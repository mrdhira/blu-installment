package constant

type VehicleType string

const (
	VTCar        VehicleType = "mobil"
	VTMotorcycle VehicleType = "motor"
)

type VehicleCondition string

const (
	VCNew    VehicleCondition = "baru"
	VCSecond VehicleCondition = "bekas"
)

const (
	CtxSQLTx = "sql_tx"
)
