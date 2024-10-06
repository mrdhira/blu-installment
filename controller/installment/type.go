package installment

import (
	"log/slog"

	"blu-installment/controller"
	"blu-installment/service"
)

type installmentCtrl struct {
	logger *slog.Logger

	installmentSvc service.IInstallmentService
}

func New(
	logger *slog.Logger,
	installmentSvc service.IInstallmentService,
) controller.IInstallmentController {
	return &installmentCtrl{logger, installmentSvc}
}
