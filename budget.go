package main

import (
	"database/sql"
	"fmt"
	"log"
)

func addBudget(
	db *sql.DB,
	q string,
	n string,
	a float64,
	id int,
	name string,
	amount float64,
) (err error) {
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
	fmt.Println(n + " budget successfully added!")
	return err
}

func getBudgets(
	db *sql.DB,
	q string,
	id int,
	name string,
	amount float64,
) (err error) {
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
	return err
}
