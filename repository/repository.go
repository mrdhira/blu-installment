package repository

import (
	"blu-installment/model"
	"context"
)

type BaseDatabaseRepository interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type IMonthlyInstallment interface {
	BaseDatabaseRepository

	Create(ctx context.Context, monthlyInstallment *model.MonthlyInstallment) error
	GetListPagination(ctx context.Context, page int, limit int) ([]*model.MonthlyInstallment, error)
	GetByRequestID(ctx context.Context, requestID string) ([]*model.MonthlyInstallment, error)
	Update(ctx context.Context, monthlyInstallment *model.MonthlyInstallment) error
	DeleteByRequestID(ctx context.Context, requestID string) error
}
