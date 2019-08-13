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
	q := "insert into user_budgets"
	n := "groceries"
	a := 250.00

	mock.ExpectQuery(q).
		WithArgs(n, a).
		WillReturnRows(sqlmock.NewRows(columns))

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

	q := "select (.+) from user_budgets"
	columns := []string{"budget_name", "allowance"}

	mock.ExpectQuery(q).
		WithArgs(1).
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

func TestShouldUpdateBudget(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	q := "update user_budgets set amount"
	a := 50.00
	columns := []string{"id", "budget_name", "allowance"}
	var (
		id     int64
		bID    int
		amount float64
	)
	id = 1
	mock.ExpectQuery(q).
		WithArgs(a, id).
		WillReturnRows(sqlmock.NewRows(columns))

	if err = updateRecord(db, q, a, id, bID, amount); err != nil {
		t.Errorf("error was not expected while updating budget: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldDeleteBudget(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	q := "delete from user_budgets"
	columns := []string{"id", "budget_name", "allowance"}
	var (
		i      int64
		id     int
		amount float64
	)
	i = 0
	mock.ExpectQuery(q).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(columns))

	if err = deleteRecord(db, q, i, id, amount); err != nil {
		t.Errorf("error was not expected while deleting budget: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldAddTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"transaction_desc", "amount_spent", "budget_id"}
	q := "insert into transactions"
	n := "dinner"
	a := 80.00
	var (
		bID      int64
		id       int
		name     string
		amount   float64
		budgetID int
	)
	bID = 1
	mock.ExpectQuery(q).
		WithArgs(n, a, bID).
		WillReturnRows(sqlmock.NewRows(columns))

	if err = addTrans(db, q, n, a, bID, id, name, amount, budgetID); err != nil {
		t.Errorf("error was not expected while adding adding a transaction: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldRetrieveAllTransactions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	q := "select (.+) from user_budgets"
	columns := []string{"transaction_desc", "amount_spent", "budget_id"}
	var (
		id       int
		name     string
		amount   float64
		budgetID int
	)
	mock.ExpectQuery(q).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns))

	if err = getTrans(db, q, id, name, amount, budgetID); err != nil {
		t.Errorf("error was not expected while adding getting all transactions: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldUpdateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	q := "update transactions set amount"
	a := 70.00
	columns := []string{"transaction_desc", "amount_spent", "budget_id"}
	var (
		id     int64
		bID    int
		amount float64
	)
	id = 2
	mock.ExpectQuery(q).
		WithArgs(a, id).
		WillReturnRows(sqlmock.NewRows(columns))

	if err = updateRecord(db, q, a, id, bID, amount); err != nil {
		t.Errorf("error was not expected while updating transaction: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldDeleteTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	q := "delete from transactions"
	columns := []string{"transaction_desc", "amount_spent", "budget_id"}
	var (
		i      int64
		id     int
		amount float64
	)
	i = 0
	mock.ExpectQuery(q).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(columns))

	if err = deleteRecord(db, q, i, id, amount); err != nil {
		t.Errorf("error was not expected while deleting transaction: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldJoinTransactionsWithBudgetsOnBudgetId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	q := "select (.+) inner join user_budgets"
	columns := []string{"transaction_id", "transaction_desc", "amount_spent", "budget_id", "budget_id", "budget_name", "allowance"}
	var (
		transactionID   int
		transactionName string
		amount          float64
		budgetID        int
		bID             int
		budgetName      string
		allowance       float64
	)
	mock.ExpectQuery(q).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns))

	if err = innerJoin(db, q, transactionID, transactionName, amount, budgetID, bID, budgetName, allowance); err != nil {
		t.Errorf("error was not expected while deleting transaction: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
