package controller

import "net/http"

type IInstallmentController interface {
	CalculateMonthlyInstallment(w http.ResponseWriter, r *http.Request)
}
