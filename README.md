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

If successful, your database will be set up.

Now, run `go build` and then any of the commands below.

## Commands

**Important:** All commands begin with `./go-budget-it`

ADD <br>
`add budget <name> <allowance>` <br>
`add transaction <name/description> <amount> <budget_id>`

VIEW <br>
`get budgets` <br>
`get transactions`

UPDATE <br>
`update budget <budget_id> <new_allowance>` <br>
`update transaction <transaction_id> <new_amount>`

DELETE <br>
`delete budget <budget_id>` <br>
`delete transaction <transaction_id>`

INNER JOIN transactions with budgets on budget_id <br>
`join-on`

## Demo

ADD
![adding budgets and transactions](./gifs/add.gif)

UPDATE
![updating budgets and transactions](./gifs/update.gif)

DELETE
![deleting budgets and transactions](./gifs/delete.gif)