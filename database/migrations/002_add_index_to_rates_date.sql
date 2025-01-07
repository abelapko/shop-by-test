-- +goose Up
CREATE INDEX idx_rates_date ON rates (date);

-- +goose Down
DROP INDEX idx_rates_date ON rates;
