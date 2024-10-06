package installment

import (
	"log/slog"

	"blu-installment/service"
)

type installmentSvc struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) service.IInstallmentService {
	return &installmentSvc{logger}
}
