package main

import (
	"database/sql"
	"fmt"
	"log"
)

func addTrans(
	db *sql.DB,
	q string,
	n string,
	a float64,
	bID int64,
	id int,
	name string,
	amount float64,
	budgetID int,
) {
	rows, err := db.Query(q, n, a, bID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &amount, &budgetID)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, amount, budgetID)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n + " transaction successfully added!")
}

func getTrans(
	db *sql.DB,
	q string,
	id int,
	name string,
	amount float64,
	budgetID int,
) {
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &amount, &budgetID)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, amount, budgetID)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
