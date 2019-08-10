DROP DATABASE IF EXISTS budget;
CREATE DATABASE budget;
USE budget;

CREATE TABLE user_budgets (
    budget_id INTEGER AUTO_INCREMENT NOT NULL,
    budget_name VARCHAR(50) NOT NULL,
    allowance DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (budget_id)
);

CREATE TABLE transactions (
    transaction_id INTEGER AUTO_INCREMENT NOT NULL,
    transaction_desc VARCHAR(100) NOT NULL,
    amount_spent DECIMAL(5,2) NOT NULL,
    budget_id INTEGER,
    PRIMARY KEY (transaction_id)
);