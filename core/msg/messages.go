package msg

import "reflect"

type Message struct {
	Key, Description string
	Type             reflect.Kind
}

// global messages
var (
	Title = Message{"title", "", reflect.Bool}
)

// site messages
var (
	BatteryConfigured = Message{"batteryConfigured", "", reflect.Bool}
	BatteryPower      = Message{"batteryPower", "", reflect.Float64}
	BatterySoC        = Message{"batterySoC", "", reflect.Bool}
	GridConfigured    = Message{"gridConfigured", "", reflect.Bool}
	GridCurrents      = Message{"gridCurrents", "", reflect.Bool}
	GridPower         = Message{"gridPower", "", reflect.Float64}
	PrioritySoC       = Message{"prioritySoC", "", reflect.Bool}
	PvConfigured      = Message{"pvConfigured", "", reflect.Bool}
	PvPower           = Message{"pvPower", "", reflect.Float64}
)

// loadpoint messages
var (
	ActivePhases          = Message{"activePhases", "", reflect.Int64}
	ChargeConfigured      = Message{"chargeConfigured", "", reflect.Bool} // remove
	ChargeCurrent         = Message{"chargeCurrent", "", reflect.Float64}
	ChargeCurrents        = Message{"chargeCurrents", "", reflect.Slice}
	ChargedEnergy         = Message{"chargedEnergy", "", reflect.Float64}
	ChargeDuration        = Message{"chargeDuration", "", reflect.Float64}
	ChargeEstimate        = Message{"chargeEstimate", "", reflect.Float64}
	ChargePower           = Message{"chargePower", "", reflect.Float64}
	ChargeRemainingEnergy = Message{"chargeRemainingEnergy", "", reflect.Float64}
	Charging              = Message{"charging", "", reflect.Bool}
	Climater              = Message{"climater", "", reflect.Float64}
	Connected             = Message{"connected", "", reflect.Bool}
	ConnectedDuration     = Message{"connectedDuration", "", reflect.Float64}
	Enabled               = Message{"enabled", "", reflect.Bool}
	HasVehicle            = Message{"hasVehicle", "", reflect.Bool}
	MaxCurrent            = Message{"maxCurrent", "", reflect.Float64}
	MinCurrent            = Message{"minCurrent", "", reflect.Float64}
	MinSoC                = Message{"minSoC", "", reflect.Float64}
	Mode                  = Message{"mode", "", reflect.String}
	Phases                = Message{"phases", "", reflect.Int64}
	Range                 = Message{"range", "", reflect.Int64}
	RemoteDisabled        = Message{"remoteDisabled", "", reflect.Bool}
	RemoteDisabledSource  = Message{"remoteDisabledSource", "", reflect.String}
	SocCapacity           = Message{"socCapacity", "", reflect.Float64}
	SocCharge             = Message{"socCharge", "", reflect.Float64}
	SocLevels             = Message{"socLevels", "", reflect.Float64}
	SocTitle              = Message{"socTitle", "", reflect.Float64}
	TargetSoC             = Message{"targetSoC", "", reflect.Float64}
	TargetTime            = Message{"targetTime", "", reflect.Struct}
	TimerActive           = Message{"timerActive", "", reflect.Bool}
	TimerSet              = Message{"timerSet", "", reflect.Bool}
)
