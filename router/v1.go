package router

import (
	"net/http"

	"blu-installment/controller"
)

func v1Group(
	installmentCtrl controller.IInstallmentController,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /installment/calculate-monthly", installmentCtrl.CalculateMonthlyInstallment)

	return mux
}
