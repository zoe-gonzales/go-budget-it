# go-budget-it
This is a CLI app that enables users to add, view, update, and delete budgets and transactions.

## Tools
Go, SQL, MySQL

## Set up
Requirements: MySQL, Go

Clone this repo.

To set up the database, enter the MySQL shell with `mysql -u username -p` and enter your password.

Run the SQl file with `path/to/project/schemas.sql`

**Note:** Above path can be obtained by running `pwd` in the directory in which the repo was cloned.

If successful, your database will be set up and you can run the commands below.

## Commands

All commands begin with `./main`

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