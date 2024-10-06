package installment

import (
	"log/slog"

	"blu-installment/service"
)

type installmentSvc struct {
	logger *slog.Logger
	// monthlyInstallmentRepo repository.IMonthlyInstallment
}

func New(
	logger *slog.Logger,
	// monthlyInstallmentRepo repository.IMonthlyInstallment,
) service.IInstallmentService {
	return &installmentSvc{
		logger,
		// monthlyInstallmentRepo,
	}
}
