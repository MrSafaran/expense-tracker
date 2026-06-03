-- +goose Up

ALTER TABLE expenses
ALTER COLUMN created_at
SET DEFAULT CURRENT_TIMESTAMP;

-- +goose Down

ALTER TABLE expenses
ALTER COLUMN created_at
DROP DEFAULT;