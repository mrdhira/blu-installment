package router

import (
	"net/http"

	"blu-installment/controller"
)

func New(
	installmentCtrl controller.IInstallmentController,
) *http.ServeMux {
	mux := http.NewServeMux()

	// v1Group is a function that returns a ServeMux with all the routes for the v1 API group
	v1GroupMux := v1Group(installmentCtrl)

	// Handle /api/v1 by stripping the prefix and forwarding the request to the v1GroupMux
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1GroupMux))

	return mux
}
