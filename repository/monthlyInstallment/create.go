package monthlyInstallment

import (
	"context"

	"blu-installment/model"
)

func (r *monthlyInstallmentRepo) Create(ctx context.Context, monthlyInstallment *model.MonthlyInstallment) error {
	sqlResult, err := r.sqlConn.ExecContext(
		ctx,
		createQuery,
		monthlyInstallment.RequestID,
		monthlyInstallment.VehicleType,
		monthlyInstallment.VehicleCondition,
		monthlyInstallment.VehicleYear,
		monthlyInstallment.TotalLoanAmount,
		monthlyInstallment.DownPayment,
		monthlyInstallment.TotalTenure,
		monthlyInstallment.Year,
		monthlyInstallment.MonthlyInstallment,
		monthlyInstallment.InterestRate,
		monthlyInstallment.CreatedAt,
		monthlyInstallment.UpdatedAt,
	)
	if err != nil {
		r.logger.ErrorContext(ctx, "Error create monthly installment", "error", err)
		return err
	}

	_, err = sqlResult.RowsAffected()
	if err != nil {
		r.logger.ErrorContext(ctx, "Error create monthly installment", "error", err)
		return err
	}

	monthlyInstallment.ID, err = sqlResult.LastInsertId()
	if err != nil {
		r.logger.ErrorContext(ctx, "Error create monthly installment", "error", err)
		return err
	}

	return nil

}
