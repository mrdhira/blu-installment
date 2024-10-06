package constant

import "errors"

// Controller Error

// Service Error
var (
	ErrVehicleTypeNewCannotBeLessThanCurrentYearMinusOne = errors.New("vehicle_condition 'baru' cannot input year less than current year - 1")
	ErrDownPaymentMustBeAtLeast35Percent                 = errors.New("vehicle_type 'baru' the down payment must be at least 35% of the total loan amount")
	ErrDownPaymentMustBeAtLeast25Percent                 = errors.New("vehicle_type 'bekas' the down payment must be at least 25% of the total loan amount")
)

var mapErrorToHTTPStatusCode = map[error]int{
	ErrVehicleTypeNewCannotBeLessThanCurrentYearMinusOne: 400,
}

// GetHTTPStatusCodeFromError is a function to get HTTP status code from service error
func GetHTTPStatusCodeFromError(err error) int {
	if code, ok := mapErrorToHTTPStatusCode[err]; ok {
		return code
	}

	return 500
}
