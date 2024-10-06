package installment

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"blu-installment/constant"
	"blu-installment/model"
)

const (
	CarBaseInterestRate        = 0.08
	MotorcycleBaseInterestRate = 0.09
	InterestOnOddYear          = 0.001
	InterestOnEvenYear         = 0.005
)

func (svc *installmentSvc) CalculateMonthlyInstallment(ctx context.Context, req *model.CalculateMonthlyInstallmentRequest) ([]*model.CalculateMonthlyInstallmentResponse, error) {
	// Vehicle condition "baru" cannot input year less than current year - 1
	if req.VehicleCondition == "baru" && req.VehicleYear < (time.Now().Year()-1) {
		return nil, constant.ErrVehicleTypeNewCannotBeLessThanCurrentYearMinusOne
	}

	// Vehicle type "baru" the down payment must be at least 35% of the total loan amount
	if req.VehicleCondition == "baru" &&
		req.DownPayment < (0.35*req.TotalLoanAmount) {
		return nil, constant.ErrDownPaymentMustBeAtLeast35Percent
	}

	// Vehicle type "bekas" the down payment must be at least 25% of the total loan amount
	if req.VehicleCondition == "bekas" &&
		req.DownPayment < (0.25*req.TotalLoanAmount) {
		return nil, constant.ErrDownPaymentMustBeAtLeast25Percent
	}

	// Calculate monthly installment
	var (
		baseInterest        float64
		principalLoanAmount float64
		resp                []*model.CalculateMonthlyInstallmentResponse
	)

	// Set Base Interest
	if req.VehicleType == "mobil" {
		baseInterest = CarBaseInterestRate
	} else {
		baseInterest = MotorcycleBaseInterestRate
	}

	// Set Principal Loan Amount
	principalLoanAmount = req.TotalLoanAmount - req.DownPayment

	for i := 1; i <= req.Tenure; i++ {
		monthlyInstallment := &model.CalculateMonthlyInstallmentResponse{}

		monthlyInstallment.Year = fmt.Sprintf("Tahun %d", i)

		totalLoanAmountPlusInterest := principalLoanAmount + (principalLoanAmount * baseInterest)
		installmentMonthLeft := 12*float64(req.Tenure) - (float64(i-1) * 12)
		mi := totalLoanAmountPlusInterest / installmentMonthLeft

		monthlyInstallment.MonthlyInstallment = fmt.Sprintf("Rp %.2f/bln", mi)
		monthlyInstallment.InterestRate = fmt.Sprintf("Suku Bunga : %.2f%%", baseInterest*100)

		// Append to response
		resp = append(resp, monthlyInstallment)

		// Calculate base interest for next year
		if i%2 == 0 {
			baseInterest += InterestOnEvenYear
		} else {
			baseInterest += InterestOnOddYear
		}

		// Set the next principal loan amount
		principalLoanAmount = totalLoanAmountPlusInterest - (mi * 12)
	}

	// For debug purpose
	respJson, _ := json.Marshal(resp)
	svc.logger.InfoContext(ctx, "monthly installment", "response", string(respJson))

	return resp, nil
}
