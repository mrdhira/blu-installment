package model

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Message string `json:"message"`
	Error   any    `json:"error"`
	Data    any    `json:"data"`
}

type CalculateMonthlyInstallmentResponse struct {
	Year               string `json:"year"`
	MonthlyInstallment string `json:"monthly_installment"`
	InterestRate       string `json:"interest_rate"`
}

func NewHttpJSONResponse(w http.ResponseWriter, statusCode int, responseMessage string, err any, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := BaseResponse{
		Message: responseMessage,
		Error:   err,
		Data:    data,
	}
	json.NewEncoder(w).Encode(resp)
}
