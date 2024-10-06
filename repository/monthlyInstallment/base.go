package monthlyInstallment

import (
	"context"
	"database/sql"

	"blu-installment/constant"
)

func (r *monthlyInstallmentRepo) Begin(ctx context.Context) (context.Context, error) {
	tx, err := r.sqlConn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, constant.CtxSQLTx, tx), nil
}

func (r *monthlyInstallmentRepo) Commit(ctx context.Context) error {
	tx, ok := ctx.Value(constant.CtxSQLTx).(*sql.Tx)
	if !ok {
		return nil
	}

	return tx.Commit()
}
func (r *monthlyInstallmentRepo) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value(constant.CtxSQLTx).(*sql.Tx)
	if !ok {
		return nil
	}

	return tx.Rollback()
}
