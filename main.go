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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	pass := os.Getenv("DB_PASS")
	dsn := "root:" + pass + "@tcp(127.0.0.1:3306)/budget"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error in db registration:", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error in db ping:", err)
	}

	var q string
	var m string

	if os.Args[1] == "budget" {
		q = "insert into user_budgets (budget_name, allowance) VALUES (?, ?)"
		m = "Budget for " + os.Args[2] + " successfully added!"
	} else if os.Args[1] == "spent" {
		q = "insert into transactions (transaction_desc, amount_spent) VALUES (?, ?)"
		m = "Transaction for " + os.Args[2] + " successfully added!"
	} else if os.Args[1] == "get" && os.Args[2] == "budgets" {
		q = "select * from user_budgets"
	} else if os.Args[1] == "get" && os.Args[2] == "transactions" {
		q = "select * from transactions"
	}

	var (
		id     int
		name   string
		amount float64
	)

	if len(os.Args) == 4 {
		a, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Error in string conversion to float:", err)
		}
		n := os.Args[2]
		rows, err := db.Query(q, n, a)
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
		fmt.Println(m)
	} else if len(os.Args) == 3 {
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
}
