-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS monthly_installments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    request_id VARCHAR(255) NOT NULL,
    vehicle_type VARCHAR(255) NOT NULL,
    vehicle_condition VARCHAR(255) NOT NULL,
    vehicle_year INT NOT NULL,
    total_loan_amount DOUBLE NOT NULL,
    down_payment DOUBLE NOT NULL,
    total_tenure INT NOT NULL,
    year INT NOT NULL,
    monthly_installment DOUBLE NOT NULL,
    interest_rate DOUBLE NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS monthly_installments;
-- +goose StatementEnd
