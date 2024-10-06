package monthlyInstallment

import (
	"blu-installment/model"
	"context"
)

func (r *monthlyInstallmentRepo) GetListPagination(ctx context.Context, page int, limit int) ([]*model.MonthlyInstallment, error) {
	var monthlyInstallments []*model.MonthlyInstallment
	sqlRows, err := r.sqlConn.QueryContext(ctx, getListPaginationQuery, (page-1)*limit, limit)
	if err != nil {
		r.logger.ErrorContext(ctx, "Error get list monthly installment", "error", err)
		return nil, err
	}
	defer sqlRows.Close()

	for sqlRows.Next() {
		monthlyInstallment := &model.MonthlyInstallment{}
		err = sqlRows.Scan(
			&monthlyInstallment.ID,
			&monthlyInstallment.RequestID,
			&monthlyInstallment.VehicleType,
			&monthlyInstallment.VehicleCondition,
			&monthlyInstallment.VehicleYear,
			&monthlyInstallment.TotalLoanAmount,
			&monthlyInstallment.DownPayment,
			&monthlyInstallment.TotalTenure,
			&monthlyInstallment.Year,
			&monthlyInstallment.MonthlyInstallment,
			&monthlyInstallment.InterestRate,
			&monthlyInstallment.CreatedAt,
			&monthlyInstallment.UpdatedAt,
		)
		if err != nil {
			r.logger.ErrorContext(ctx, "Error scan monthly installment", "error", err)
			return nil, err
		}
		monthlyInstallments = append(monthlyInstallments, monthlyInstallment)
	}

	return monthlyInstallments, nil
}

func (r *monthlyInstallmentRepo) GetByRequestID(ctx context.Context, requestID string) ([]*model.MonthlyInstallment, error) {
	var monthlyInstallments []*model.MonthlyInstallment

	sqlRows, err := r.sqlConn.QueryContext(ctx, getByRequestIDQuery, requestID)
	if err != nil {
		r.logger.ErrorContext(ctx, "Error get monthly installment by request id", "error", err)
		return nil, err
	}
	defer sqlRows.Close()

	for sqlRows.Next() {
		monthlyInstallment := &model.MonthlyInstallment{}
		err = sqlRows.Scan(
			&monthlyInstallment.ID,
			&monthlyInstallment.RequestID,
			&monthlyInstallment.VehicleType,
			&monthlyInstallment.VehicleCondition,
			&monthlyInstallment.VehicleYear,
			&monthlyInstallment.TotalLoanAmount,
			&monthlyInstallment.DownPayment,
			&monthlyInstallment.TotalTenure,
			&monthlyInstallment.Year,
			&monthlyInstallment.MonthlyInstallment,
			&monthlyInstallment.InterestRate,
			&monthlyInstallment.CreatedAt,
			&monthlyInstallment.UpdatedAt,
		)
		if err != nil {
			r.logger.ErrorContext(ctx, "Error scan monthly installment", "error", err)
			return nil, err
		}
		monthlyInstallments = append(monthlyInstallments, monthlyInstallment)
	}

	return monthlyInstallments, nil
}
