package installment

import (
	"context"
	"log/slog"
	"os"
	"reflect"
	"testing"
	"time"

	"blu-installment/constant"
	"blu-installment/model"
)

func TestCalculateMonthlyInstallment(t *testing.T) {
	testCases := []struct {
		name         string
		request      *model.CalculateMonthlyInstallmentRequest
		expectedResp []*model.CalculateMonthlyInstallmentResponse
		expectedErr  error
	}{
		{
			name: "Error: ErrVehicleTypeNewCannotBeLessThanCurrentYearMinusOne",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "baru",
				VehicleYear:      time.Now().Year() - 2,
				DownPayment:      100000000,
				TotalLoanAmount:  1000000000,
				Tenure:           1,
			},
			expectedResp: nil,
			expectedErr:  constant.ErrVehicleTypeNewCannotBeLessThanCurrentYearMinusOne,
		},
		{
			name: "Error: ErrDownPaymentMustBeAtLeast35Percent",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "baru",
				VehicleYear:      time.Now().Year() - 1,
				DownPayment:      100000000,
				TotalLoanAmount:  1000000000,
				Tenure:           1,
			},
			expectedResp: nil,
			expectedErr:  constant.ErrDownPaymentMustBeAtLeast35Percent,
		},
		{
			name: "Error: ErrDownPaymentMustBeAtLeast25Percent",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "bekas",
				VehicleYear:      2020,
				DownPayment:      100000000,
				TotalLoanAmount:  1000000000,
				Tenure:           1,
			},
			expectedResp: nil,
			expectedErr:  constant.ErrDownPaymentMustBeAtLeast25Percent,
		},
		{
			name: "Success: Test Cases From Sheets Given By Interviewer",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "bekas",
				VehicleYear:      2016,
				TotalLoanAmount:  100000000,
				DownPayment:      25000000,
				Tenure:           3,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 2250000.00/bln",
					InterestRate:       "Suku Bunga : 8.00%",
				},
				{
					Year:               "Tahun 2",
					MonthlyInstallment: "Rp 2432250.00/bln",
					InterestRate:       "Suku Bunga : 8.10%",
				},
				{
					Year:               "Tahun 3",
					MonthlyInstallment: "Rp 2641423.50/bln",
					InterestRate:       "Suku Bunga : 8.60%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Mobil Baru 1 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "baru",
				VehicleYear:      time.Now().Year() - 1,
				TotalLoanAmount:  1000000000,
				DownPayment:      350000000,
				Tenure:           1,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 58500000.00/bln",
					InterestRate:       "Suku Bunga : 8.00%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Mobil Baru 6 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "baru",
				VehicleYear:      time.Now().Year() - 1,
				TotalLoanAmount:  1000000000,
				DownPayment:      350000000,
				Tenure:           6,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 9750000.00/bln",
					InterestRate:       "Suku Bunga : 8.00%",
				},
				{
					Year:               "Tahun 2",
					MonthlyInstallment: "Rp 10539750.00/bln",
					InterestRate:       "Suku Bunga : 8.10%",
				},
				{
					Year:               "Tahun 3",
					MonthlyInstallment: "Rp 11446168.50/bln",
					InterestRate:       "Suku Bunga : 8.60%",
				},
				{
					Year:               "Tahun 4",
					MonthlyInstallment: "Rp 12441985.16/bln",
					InterestRate:       "Suku Bunga : 8.70%",
				},
				{
					Year:               "Tahun 5",
					MonthlyInstallment: "Rp 13586647.79/bln",
					InterestRate:       "Suku Bunga : 9.20%",
				},
				{
					Year:               "Tahun 6",
					MonthlyInstallment: "Rp 14850206.04/bln",
					InterestRate:       "Suku Bunga : 9.30%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Mobil Bekas 1 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "bekas",
				VehicleYear:      2020,
				TotalLoanAmount:  1000000000,
				DownPayment:      250000000,
				Tenure:           1,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 67500000.00/bln",
					InterestRate:       "Suku Bunga : 8.00%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Mobil Bekas 6 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "mobil",
				VehicleCondition: "bekas",
				VehicleYear:      2020,
				TotalLoanAmount:  1000000000,
				DownPayment:      250000000,
				Tenure:           6,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 11250000.00/bln",
					InterestRate:       "Suku Bunga : 8.00%",
				},
				{
					Year:               "Tahun 2",
					MonthlyInstallment: "Rp 12161250.00/bln",
					InterestRate:       "Suku Bunga : 8.10%",
				},
				{
					Year:               "Tahun 3",
					MonthlyInstallment: "Rp 13207117.50/bln",
					InterestRate:       "Suku Bunga : 8.60%",
				},
				{
					Year:               "Tahun 4",
					MonthlyInstallment: "Rp 14356136.72/bln",
					InterestRate:       "Suku Bunga : 8.70%",
				},
				{
					Year:               "Tahun 5",
					MonthlyInstallment: "Rp 15676901.30/bln",
					InterestRate:       "Suku Bunga : 9.20%",
				},
				{
					Year:               "Tahun 6",
					MonthlyInstallment: "Rp 17134853.12/bln",
					InterestRate:       "Suku Bunga : 9.30%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Motor Baru 1 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "motor",
				VehicleCondition: "baru",
				VehicleYear:      time.Now().Year() - 1,
				TotalLoanAmount:  1000000000,
				DownPayment:      350000000,
				Tenure:           1,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 59041666.67/bln",
					InterestRate:       "Suku Bunga : 9.00%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Motor Baru 6 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "motor",
				VehicleCondition: "baru",
				VehicleYear:      time.Now().Year() - 1,
				TotalLoanAmount:  1000000000,
				DownPayment:      350000000,
				Tenure:           6,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 9840277.78/bln",
					InterestRate:       "Suku Bunga : 9.00%",
				},
				{
					Year:               "Tahun 2",
					MonthlyInstallment: "Rp 10735743.06/bln",
					InterestRate:       "Suku Bunga : 9.10%",
				},
				{
					Year:               "Tahun 3",
					MonthlyInstallment: "Rp 11766374.39/bln",
					InterestRate:       "Suku Bunga : 9.60%",
				},
				{
					Year:               "Tahun 4",
					MonthlyInstallment: "Rp 12907712.70/bln",
					InterestRate:       "Suku Bunga : 9.70%",
				},
				{
					Year:               "Tahun 5",
					MonthlyInstallment: "Rp 14224299.40/bln",
					InterestRate:       "Suku Bunga : 10.20%",
				},
				{
					Year:               "Tahun 6",
					MonthlyInstallment: "Rp 15689402.24/bln",
					InterestRate:       "Suku Bunga : 10.30%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Motor Bekas 1 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "motor",
				VehicleCondition: "bekas",
				VehicleYear:      2020,
				TotalLoanAmount:  1000000000,
				DownPayment:      250000000,
				Tenure:           1,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 68125000.00/bln",
					InterestRate:       "Suku Bunga : 9.00%",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Success: Motor Bekas 6 year",
			request: &model.CalculateMonthlyInstallmentRequest{
				VehicleType:      "motor",
				VehicleCondition: "bekas",
				VehicleYear:      2020,
				TotalLoanAmount:  1000000000,
				DownPayment:      250000000,
				Tenure:           6,
			},
			expectedResp: []*model.CalculateMonthlyInstallmentResponse{
				{
					Year:               "Tahun 1",
					MonthlyInstallment: "Rp 11354166.67/bln",
					InterestRate:       "Suku Bunga : 9.00%",
				},
				{
					Year:               "Tahun 2",
					MonthlyInstallment: "Rp 12387395.83/bln",
					InterestRate:       "Suku Bunga : 9.10%",
				},
				{
					Year:               "Tahun 3",
					MonthlyInstallment: "Rp 13576585.83/bln",
					InterestRate:       "Suku Bunga : 9.60%",
				},
				{
					Year:               "Tahun 4",
					MonthlyInstallment: "Rp 14893514.66/bln",
					InterestRate:       "Suku Bunga : 9.70%",
				},
				{
					Year:               "Tahun 5",
					MonthlyInstallment: "Rp 16412653.15/bln",
					InterestRate:       "Suku Bunga : 10.20%",
				},
				{
					Year:               "Tahun 6",
					MonthlyInstallment: "Rp 18103156.43/bln",
					InterestRate:       "Suku Bunga : 10.30%",
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

			svc := New(logger)

			resp, err := svc.CalculateMonthlyInstallment(context.Background(), tc.request)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(tc.expectedResp, resp) {
				t.Errorf("Expected response %v, got %v", tc.expectedResp, resp)
			}
		})
	}
}
