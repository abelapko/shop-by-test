-- +goose Up
CREATE TABLE rates (
  id INT AUTO_INCREMENT PRIMARY KEY,
  code VARCHAR(10),
  nominal INT,
  name VARCHAR(100),
  rate DECIMAL(10, 4),
  date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE rates;
