DROP DATABASE currency_api;
CREATE DATABASE currency_api;
USE currency_api;

SHOW TABLES;
SELECT * FROM currency;
SELECT currency from currency;

CREATE TABLE currency (
    id INT NOT NULL PRIMARY KEY auto_increment,
    currency VARCHAR(20) NOT NULL,
    currency_value DECIMAL(8,2) NOT NULL
);

INSERT INTO currency (currency, currency_value) VALUES ("usd", 5.57);
INSERT INTO currency (currency, currency_value) VALUES ("eur", 6.34);
INSERT INTO currency (currency, currency_value) VALUES ("gbp", 7.54);