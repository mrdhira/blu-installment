package monthlyInstallment

import "context"

func (r *monthlyInstallmentRepo) DeleteByRequestID(ctx context.Context, requestID string) error {
	sqlResult, err := r.sqlConn.ExecContext(
		ctx,
		deleteByRequestIDQuery,
		requestID,
	)
	if err != nil {
		r.logger.ErrorContext(ctx, "Error delete monthly installment", "error", err)
		return err
	}

	_, err = sqlResult.RowsAffected()
	if err != nil {
		r.logger.ErrorContext(ctx, "Error delete monthly installment", "error", err)
		return err
	}

	return nil
}
