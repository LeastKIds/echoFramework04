CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_up TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
