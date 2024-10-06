package model

import "time"

type MonthlyInstallment struct {
	ID                 int64      `db:"id" json:"id"`
	RequestID          string     `db:"request_id" json:"request_id"`
	VehicleType        string     `db:"vehicle_type" json:"vehicle_type"`
	VehicleCondition   string     `db:"vehicle_condition" json:"vehicle_condition"`
	VehicleYear        int        `db:"vehicle_year" json:"vehicle_year"`
	TotalLoanAmount    float64    `db:"total_loan_amount" json:"total_loan_amount"`
	DownPayment        float64    `db:"down_payment" json:"down_payment"`
	TotalTenure        int        `db:"total_tenure" json:"total_tenure"`
	Year               int        `db:"year" json:"year"`
	MonthlyInstallment float64    `db:"monthly_installment" json:"monthly_installment"`
	InterestRate       float64    `db:"interest_rate" json:"interest_rate"`
	CreatedAt          time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt          *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (m MonthlyInstallment) TableName() string {
	return "monthly_installments"
}
