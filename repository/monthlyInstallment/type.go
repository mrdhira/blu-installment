package monthlyInstallment

import (
	"blu-installment/repository"
	"database/sql"
	"log/slog"
)

type monthlyInstallmentRepo struct {
	logger  *slog.Logger
	sqlConn *sql.DB
}

var (
	createQuery            = `INSERT INTO monthly_installments (request_id, vehicle_type, vehicle_condition, vehicle_year, total_loan_amount, down_payment, total_tenure, year, monthly_installment, interest_rate, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	getListPaginationQuery = `SELECT id, request_id, vehicle_type, vehicle_condition, vehicle_year, total_loan_amount, down_payment, total_tenure, year, monthly_installment, interest_rate, created_at, updated_at FROM monthly_installments WHERE deleted_at IS NULL ORDER BY created_at DESC, request_id ASC, tenure ASC DESC LIMIT ? OFFSET ?`
	getByRequestIDQuery    = `SELECT id, request_id, vehicle_type, vehicle_condition, vehicle_year, total_loan_amount, down_payment, total_tenure, year, monthly_installment, interest_rate, created_at, updated_at FROM monthly_installments WHERE request_id = ? AND deleted_at IS NULL`
	updateQuery            = `UPDATE monthly_installments SET total_loan_amount = ?, down_payment = ?, total_tenure = ?, year = ?, monthly_installment = ?, interest_rate = ?, created_at = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	deleteByRequestIDQuery = `UPDATE monthly_installments SET deleted_at = NOW() WHERE request_id = $1`
)

func New(logger *slog.Logger, sqlConn *sql.DB) repository.IMonthlyInstallment {
	return &monthlyInstallmentRepo{
		logger:  logger,
		sqlConn: sqlConn,
	}
}
