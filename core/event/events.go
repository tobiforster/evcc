package event

const (
	ChargeStart       = "start"      // update chargeTimer
	ChargeStop        = "stop"       // update chargeTimer
	ChargeCurrent     = "current"    // update fakeChargeMeter
	ChargePower       = "power"      // update chargeRater
	VehicleConnect    = "connect"    // vehicle connected
	VehicleDisconnect = "disconnect" // vehicle disconnected
)
