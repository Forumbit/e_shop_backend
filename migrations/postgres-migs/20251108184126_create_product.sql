-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY
    ,title VARCHAR(255) NOT NULL
    ,description TEXT NOT NULL
    ,rating NUMERIC(2, 1) CHECK (rating >= 0 AND rating <= 5)
    ,category_id INTEGER REFERENCES categories(id)
    ,brand_id INTEGER REFERENCES brands(id)
    ,price NUMERIC(10, 2) NOT NULL
    ,discount NUMERIC(5, 2) CHECK (discount >= 0 AND discount <= 100)
    ,tags TEXT[] NOT NULL
    ,weight INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
