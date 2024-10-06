package model

type CalculateMonthlyInstallmentRequest struct {
	VehicleType      string  `json:"vehicle_type"`
	VehicleCondition string  `json:"vehicle_condition"`
	VehicleYear      int     `json:"vehicle_year"`
	TotalLoanAmount  float64 `json:"total_loan_amount"`
	DownPayment      float64 `json:"down_payment"`
	Tenure           int     `json:"tenure"`
}
