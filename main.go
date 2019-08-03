package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
	fmt.Println("All is running as it should.....")
}
