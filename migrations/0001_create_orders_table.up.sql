CREATE TABLE IF NOT EXISTS orders_db (
    id VARCHAR(255) PRIMARY KEY,
    product VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    total_price FLOAT NOT NULL
    );
