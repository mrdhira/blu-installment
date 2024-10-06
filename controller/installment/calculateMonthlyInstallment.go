package installment

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"blu-installment/constant"
	"blu-installment/model"
)

func (ctrl *installmentCtrl) validateCalculateMonthlyInstallmentRequest(req model.CalculateMonthlyInstallmentRequest) error {
	req.VehicleType = strings.ToLower(req.VehicleType)
	req.VehicleCondition = strings.ToLower(req.VehicleCondition)

	// Check if the vehicle type is either "mobil" or "motor" with case insensitive
	// If not return error
	if !strings.Contains(req.VehicleType, string(constant.VTCar)) &&
		!strings.Contains(req.VehicleType, string(constant.VTMotorcycle)) {
		return errors.New("vehicle_type must be between 'mobil' or 'motor'")
	}

	// Check if the vehicle condition is either "baru" or "bekas" with case insensitive
	// If not return error
	if !strings.Contains(req.VehicleCondition, string(constant.VCNew)) &&
		!strings.Contains(req.VehicleCondition, string(constant.VCSecond)) {
		return errors.New("vehicle_condition must be between 'baru' or 'bekas'")
	}

	// Check if the vehicle year is 4 digit and between 1998 and current year
	// If not return error
	if req.VehicleYear < 1998 || req.VehicleYear > time.Now().Year() {
		return errors.New("vehicle_year must be 4 digit and between 1998 and current year")
	}

	// Check if the Total Loan Amount is greater than 1 billion rupiah
	// If not return error
	if req.TotalLoanAmount > 1000000000 {
		return errors.New("vehicle_price must be less than 1 billion rupiah")
	}

	// Check if the tenure is between 1 and 6 years
	// If not return error
	if req.Tenure < 1 || req.Tenure > 6 {
		return errors.New("tenure must be between 1 and 6 years")
	}

	return nil
}

func (ctrl *installmentCtrl) CalculateMonthlyInstallment(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	// Decode request
	var req model.CalculateMonthlyInstallmentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		ctrl.logger.ErrorContext(ctx, "failed to decode request", "error", err)
		model.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to decode request", err, nil)
		return
	}

	// Validate request
	err = ctrl.validateCalculateMonthlyInstallmentRequest(req)
	if err != nil {
		ctrl.logger.ErrorContext(ctx, "failed to validate request", "error", err)
		model.NewHttpJSONResponse(w, http.StatusBadRequest, "failed to validate request", err.Error(), nil)
		return
	}

	// Calculate monthly installment
	resp, err := ctrl.installmentSvc.CalculateMonthlyInstallment(ctx, &req)
	if err != nil {
		ctrl.logger.ErrorContext(ctx, "failed to calculate monthly installment", "error", err)
		model.NewHttpJSONResponse(w, constant.GetHTTPStatusCodeFromError(err), "failed to calculate monthly installment", err.Error(), nil)
		return
	}

	model.NewHttpJSONResponse(w, http.StatusOK, "success", nil, resp)
}
