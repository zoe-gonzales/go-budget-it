# go-budget-it
This is a CLI app that enables users to add, view, update, and delete budgets and transactions.

## Tools
Go, SQL, MySQL

## Set up
**Requirements:** MySQL, Go

Clone this repo.

To set up the database, enter the MySQL shell with `mysql -u username -p` and enter your password.

Create database and tables: `source path/to/project/schemas.sql`

**Note:** Above path can be obtained by running `pwd` in the directory in which the repo was cloned.

If successful, your database will be set up and no errors will print to the console when `go run main.go` is run.

Now, run `go build main.go` and then any of the commands below.

## Commands

**Important:** All commands begin with `./main`

ADD
`add budget <name> <allowance>`
`add transaction <name/description> <amount> <budget_id>`

VIEW
`get budgets`
`get transactions`

UPDATE
`update budget <id> <allowance>`
`update transaction <id> <amount>`

DELETE
`delete budget <id>`
`delete transaction <id>`