package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type query struct {
	command  string
	table    string
	name     string
	amount   float64
	id       int64
	budgetID int64
}

func main() {
	// loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	pass := os.Getenv("DB_PASS")
	dsn := "root:" + pass + "@tcp(127.0.0.1:3306)/budget"

	// establishing reference to db
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error in db registration:", err)
		os.Exit(1)
	}
	// checking access to db
	err = db.Ping()
	if err != nil {
		fmt.Println("Error in db ping:", err)
	}

	if len(os.Args) == 1 {
		fmt.Println("No commands entered.")
		fmt.Println("See readme at https://github.com/zoe-gonzales/go-budget-it for all commands.")
		os.Exit(1)
	}

	newQuery := query{
		command: os.Args[1],
	}

	if len(os.Args) > 2 {
		newQuery.table = os.Args[2]
	}

	// setting name and amount for adding a budget/transaction
	if newQuery.command == "add" {
		a, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			fmt.Println("Error in amount conversion to float:", err)
		}
		if len(os.Args) > 5 {
			budgetID, err := strconv.ParseInt(os.Args[5], 10, 64)
			if err != nil {
				fmt.Println("Error in id conversion to int:", err)
			}
			newQuery.budgetID = budgetID
		}
		newQuery.name = os.Args[3]
		newQuery.amount = a
	}
	// setting id for delete and update actions
	if len(os.Args) >= 4 && newQuery.command != "add" {
		id, err := strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil {
			fmt.Println("Error in id conversion to integer:", err)
		}
		newQuery.id = id
	}
	// setting amount for update action
	if len(os.Args) >= 5 && newQuery.command != "add" {
		a, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			fmt.Println("Error in amount conversion to float:", err)
		}
		newQuery.amount = a
	}

	var (
		q               string
		id              int
		name            string
		amount          float64
		budgetID        int
		transactionID   int
		budgetName      string
		transactionName string
		allowance       float64
	)

	// Insert a new budget
	if newQuery.command == "add" && newQuery.table == "budget" {
		q = "insert into user_budgets (budget_name, allowance) VALUES (?, ?)"
		addBudget(db, q, newQuery.name, newQuery.amount, id, name, amount)
	}
	// Insert a new transaction
	if newQuery.command == "add" && newQuery.table == "transaction" {
		q = "insert into transactions (transaction_desc, amount_spent, budget_id) VALUES (?, ?, ?)"
		addTrans(db, q, newQuery.name, newQuery.amount, newQuery.budgetID, id, name, amount, budgetID)
	}
	// Select all records from budgets
	if newQuery.command == "get" && newQuery.table == "budgets" {
		q = "select * from user_budgets"
		getBudgets(db, q, id, name, amount)
	}
	// Select all records from transactions
	if newQuery.command == "get" && newQuery.table == "transactions" {
		q = "select * from transactions"
		getTrans(db, q, id, name, amount, budgetID)
	}
	// update a budget
	if newQuery.command == "update" && newQuery.table == "budget" {
		q = "update user_budgets set allowance = ? where budget_id = ?"
		updateRecord(db, q, newQuery.amount, newQuery.id, id, amount)
	}
	// update a transaction
	if newQuery.command == "update" && newQuery.table == "transaction" {
		q = "update transactions set amount_spent = ? where transaction_id = ?"
		updateRecord(db, q, newQuery.amount, newQuery.id, id, amount)
	}
	// delete a budget
	if newQuery.command == "delete" && newQuery.table == "budget" {
		q = "delete from user_budgets where budget_id = ?"
		deleteRecord(db, q, newQuery.id, id, amount)
	}
	// delete a transaction
	if newQuery.command == "delete" && newQuery.table == "transaction" {
		q = "delete from transactions where transaction_id = ?"
		deleteRecord(db, q, newQuery.id, id, amount)
	}
	// inner join transactions with budgets on budget_id
	if newQuery.command == "join-on" {
		q = "select * from transactions inner join user_budgets on transactions.budget_id=user_budgets.budget_id"
		innerJoin(db, q, transactionID, transactionName, amount, budgetID, budgetID, budgetName, allowance)
	}
}
