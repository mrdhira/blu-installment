package monthlyInstallment

import (
	"blu-installment/model"
	"context"
)

func (r *monthlyInstallmentRepo) Update(ctx context.Context, monthlyInstallment *model.MonthlyInstallment) error {
	sqlResult, err := r.sqlConn.ExecContext(
		ctx,
		updateQuery,
		monthlyInstallment.VehicleType,
		monthlyInstallment.VehicleCondition,
		monthlyInstallment.VehicleYear,
		monthlyInstallment.TotalLoanAmount,
		monthlyInstallment.DownPayment,
		monthlyInstallment.TotalTenure,
		monthlyInstallment.Year,
		monthlyInstallment.MonthlyInstallment,
		monthlyInstallment.InterestRate,
		monthlyInstallment.UpdatedAt,
		monthlyInstallment.ID,
	)
	if err != nil {
		r.logger.ErrorContext(ctx, "Error update monthly installment", "error", err)
		return err
	}

	_, err = sqlResult.RowsAffected()
	if err != nil {
		r.logger.ErrorContext(ctx, "Error update monthly installment", "error", err)
		return err
	}

	return nil
}
