package service

import (
	"context"

	"blu-installment/model"
)

type IInstallmentService interface {
	CalculateMonthlyInstallment(ctx context.Context, req *model.CalculateMonthlyInstallmentRequest) ([]*model.CalculateMonthlyInstallmentResponse, error)
}
