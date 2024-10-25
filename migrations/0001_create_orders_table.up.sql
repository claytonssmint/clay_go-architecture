CREATE TABLE IF NOT EXISTS produtos (
    id VARCHAR(255) PRIMARY KEY,
    product VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    total_price FLOAT NOT NULL
    );

SHOW DATABASES;
CREATE DATABASE IF NOT EXISTS pedidos_db;
SHOW TABLES;
USE pedidos_db;
DESCRIBE produtos;
SELECT * FROM produtos;