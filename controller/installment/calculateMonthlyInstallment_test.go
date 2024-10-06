package installment

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"blu-installment/model"
	installmentService "blu-installment/service/installment"
)

func TestCalculateMonthlyInstallment(t *testing.T) {
	// Init Test
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	installmentSvc := installmentService.New(
		logger,
	)
	installmentCtrl := New(logger, installmentSvc)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/installment/calculate-monthly-installment", installmentCtrl.CalculateMonthlyInstallment)

	respSuccessFromTestCaseByInterviewer := model.BaseResponse{
		Message: "success",
		Error:   nil,
		Data: []*model.CalculateMonthlyInstallmentResponse{
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
	}
	respSuccessFromTestCaseByInterviewerJSON, err := json.Marshal(respSuccessFromTestCaseByInterviewer)
	if err != nil {
		t.Errorf("failed to marshal response: %v", err)
	}

	testCases := []struct {
		name                 string
		requestPayload       string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Error: failed to decode request",
			requestPayload:       `{"vehicle_year": "this is error"}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"failed to decode request","error":{"Value":"string","Type":{},"Offset":32,"Struct":"CalculateMonthlyInstallmentRequest","Field":"vehicle_year"},"data":null}`,
		},
		{
			name:                 "Error: failed to validate request: vehicle_type must be between 'mobil' or 'motor'",
			requestPayload:       `{"vehicle_type": "this is error"}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"failed to validate request","error":"vehicle_type must be between 'mobil' or 'motor'","data":null}`,
		},
		{
			name:                 "Error: failed to validate request: vehicle_condition must be between 'baru' or 'bekas'",
			requestPayload:       `{"vehicle_type": "mobil", "vehicle_condition": "this is error"}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"failed to validate request","error":"vehicle_condition must be between 'baru' or 'bekas'","data":null}`,
		},
		{
			name:                 "Error: failed to validate request: vehicle_year must be 4 digit and between 1998 and current year",
			requestPayload:       `{"vehicle_type": "mobil", "vehicle_condition": "baru", "vehicle_year": 1997}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"failed to validate request","error":"vehicle_year must be 4 digit and between 1998 and current year","data":null}`,
		},
		{
			name:                 "Error: failed to validate request: vehicle_price must be less than 1 billion rupiah",
			requestPayload:       `{"vehicle_type": "mobil", "vehicle_condition": "baru", "vehicle_year": 2021, "total_loan_amount": 10000000000}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"failed to validate request","error":"vehicle_price must be less than 1 billion rupiah","data":null}`,
		},
		{
			name:                 "Error: failed to validate request: tenure must be between 1 and 6 years",
			requestPayload:       `{"vehicle_type": "mobil", "vehicle_condition": "baru", "vehicle_year": 2021, "total_loan_amount": 100000000, "tenure": 7}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"failed to validate request","error":"tenure must be between 1 and 6 years","data":null}`,
		},
		{
			name:                 "Error: failed to calculate monthly installment",
			requestPayload:       `{"vehicle_type": "mobil", "vehicle_condition": "baru", "vehicle_year": 2021, "total_loan_amount": 100000000, "down_payment": 10000000, "tenure": 6}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"failed to calculate monthly installment","error":"vehicle_condition 'baru' cannot input year less than current year - 1","data":null}`,
		},
		{
			name:                 "Success: Test Cases From Sheets Given By Interviewer",
			requestPayload:       `{"vehicle_type": "mobil", "vehicle_condition": "bekas", "vehicle_year": 2016, "total_loan_amount": 100000000, "down_payment": 25000000, "tenure": 3}`,
			expectedStatusCode:   200,
			expectedResponseBody: string(respSuccessFromTestCaseByInterviewerJSON),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/api/v1/installment/calculate-monthly-installment", strings.NewReader(tc.requestPayload))

			mux.ServeHTTP(recorder, request)
			if recorder.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, recorder.Code)
			}

			var expectedResponseBodyMap, actualResponseBodyMap map[string]interface{}

			// Unmarshal expected response body
			err := json.Unmarshal([]byte(tc.expectedResponseBody), &expectedResponseBodyMap)
			if err != nil {
				t.Errorf("failed to unmarshal expected response body: %v", err)
			}

			// Unmarshal actual response body
			err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBodyMap)
			if err != nil {
				t.Errorf("failed to unmarshal actual response body: %v", err)
			}

			// Compare response body
			if !reflect.DeepEqual(expectedResponseBodyMap, actualResponseBodyMap) {
				t.Errorf("expected response body %v, got %v", expectedResponseBodyMap, actualResponseBodyMap)
			}
		})
	}
}
