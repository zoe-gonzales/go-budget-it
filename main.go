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
	command string
	table   string
	name    string
	amount  float64
	id      int64
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

	var q string

	newQuery := query{
		command: os.Args[1],
		table:   os.Args[2],
	}

	// setting name and amount for adding a budget/transaction
	if newQuery.command == "add" {
		a, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			fmt.Println("Error in amount conversion to float:", err)
		}
		newQuery.name = os.Args[3]
		newQuery.amount = a
	}
	// setting id for delete and update actions
	if len(os.Args) >= 4 {
		id, err := strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil {
			fmt.Println("Error in id conversion to integer:", err)
		}
		newQuery.id = id
	}
	// setting amount for update action
	if len(os.Args) >= 5 {
		a, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			fmt.Println("Error in amount conversion to float:", err)
		}
		newQuery.amount = a
	}

	// QUERIES

	// add budget
	if newQuery.command == "add" && newQuery.table == "budget" {
		q = "insert into user_budgets (budget_name, allowance) VALUES (?, ?)"
	}
	// add transaction
	if newQuery.command == "add" && newQuery.table == "transaction" {
		q = "insert into transactions (transaction_desc, amount_spent) VALUES (?, ?)"
	}
	// get all budgets
	if newQuery.command == "get" && newQuery.table == "budgets" {
		q = "select * from user_budgets"
	}
	// get all transactions
	if newQuery.command == "get" && newQuery.table == "transactions" {
		q = "select * from transactions"
	}
	// update a budget
	if newQuery.command == "update" && newQuery.table == "budget" {
		q = "update user_budgets set allowance = ? where budget_id = ?"
	}
	// update a transaction
	if newQuery.command == "update" && newQuery.table == "transaction" {
		q = "update transactions set amount_spent = ? where transaction_id = ?"
	}
	// delete a budget
	if newQuery.command == "delete" && newQuery.table == "budget" {
		q = "delete from user_budgets where budget_id = ?"
	}
	// delete a transaction
	if newQuery.command == "delete" && newQuery.table == "transaction" {
		q = "delete from transactions where transaction_id = ?"
	}

	var (
		id     int
		name   string
		amount float64
	)

	// CRUD actions

	// Update existing budget/transaction
	if newQuery.command == "update" {
		rows, err := db.Query(
			q,
			newQuery.amount,
			newQuery.id,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &amount)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, amount)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(newQuery.table + " successfully updated!")
	}

	// Insert a new budget/transaction
	if newQuery.command == "add" {
		rows, err := db.Query(
			q,
			newQuery.name,
			newQuery.amount,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name, &amount)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name, amount)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(newQuery.name + " " + newQuery.table + " successfully added!")
	}

	// Select all records from budgets/transactions
	if newQuery.command == "get" {
		rows, err := db.Query(q)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name, &amount)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name, amount)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Delete a budget/transaction
	if newQuery.command == "delete" {
		rows, err := db.Query(
			q,
			newQuery.id,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &amount)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, amount)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(newQuery.table + " successfully deleted.")
	}
}
