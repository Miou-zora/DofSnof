CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price1 DECIMAL(10, 0) NOT NULL,
    price10 DECIMAL(10, 0) NOT NULL,
    price100 DECIMAL(10, 0) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);