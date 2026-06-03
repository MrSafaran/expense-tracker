-- +goose Up

CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    amount NUMERIC NOT NULL,
    category TEXT NOT NULL
);

-- +goose Down

DROP TABLE expenses;