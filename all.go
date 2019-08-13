package main

import (
	"database/sql"
	"fmt"
	"log"
)

func updateRecord(
	db *sql.DB,
	q string,
	a float64,
	i int64,
	id int,
	amount float64,
) (err error) {
	rows, err := db.Query(q, a, i)
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
	fmt.Println("Record successfully updated!")
	return err
}

func deleteRecord(
	db *sql.DB,
	q string,
	i int64,
	id int,
	amount float64,
) (err error) {
	rows, err := db.Query(q, i)
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
	fmt.Println("Record successfully deleted.")
	return err
}

func innerJoin(
	db *sql.DB,
	q string,
	transactionID int,
	transactionName string,
	amount float64,
	budgetID int,
	bID int,
	budgetName string,
	allowance float64,
) (err error) {
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&transactionID,
			&transactionName,
			&amount,
			&budgetID,
			&bID,
			&budgetName,
			&allowance,
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(
			transactionID,
			transactionName,
			amount,
			budgetID,
			bID,
			budgetName,
			allowance,
		)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
