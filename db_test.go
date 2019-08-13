package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldAddBudget(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"o_budget_name", "o_allowance"}
	q := "insert into user_budgets\\(budget_name, allowance\\) values\\(\\?, \\?\\)"
	n := "groceries"
	a := 250.00

	mock.ExpectQuery(q).
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows(columns))
	mock.ExpectCommit()

	var (
		id     int
		name   string
		amount float64
	)

	if err = addBudget(db, q, n, a, id, name, amount); err != nil {
		t.Errorf("error was not expected while adding adding a budget: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldRetrieveAllBudgets(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	q := "select * from user_budgets"
	columns := []string{"id", "budget_name", "allowance"}

	mock.ExpectQuery(q).
		WithArgs(0).
		WillReturnRows(sqlmock.NewRows(columns))
	mock.ExpectCommit()

	var (
		id     int
		name   string
		amount float64
	)

	if err = getBudgets(db, q, id, name, amount); err != nil {
		t.Errorf("error was not expected while getting budgets: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
