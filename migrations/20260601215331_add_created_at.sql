-- +goose Up
ALTER TABLE expenses
ADD COLUMN created_at TIMESTAMP;

-- +goose Down
ALTER TABLE expenses
DROP COLUMN created_at;