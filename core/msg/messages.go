package msg

import "strings"

//go:generate enumer -type=Message

type Message int

func (m Message) Key() string {
	s := m.String()
	return strings.ToLower(string(s[0])) + s[1:]
}

const (
	_ Message = iota

	// global messages
	Error
	Warn
	AvailableVersion
	ReleaseNotes

	// site messages
	BatteryConfigured
	BatteryPower
	BatterySoC
	GridConfigured
	GridCurrents
	GridPower
	PrioritySoC
	PvConfigured
	PvPower
	SiteTitle

	// loadpoint messages
	ActivePhases
	ChargeConfigured
	ChargeCurrent
	ChargeCurrents
	ChargedEnergy
	ChargeDuration
	ChargeEstimate
	ChargePower
	ChargeRemainingEnergy
	Charging
	Climater
	Connected
	ConnectedDuration
	Enabled
	HasVehicle
	MaxCurrent
	MinCurrent
	MinSoC
	Mode
	Phases
	Range
	RemoteDisabled
	RemoteDisabledSource
	SocCapacity
	SocCharge
	SocLevels
	SocTitle
	TargetSoC
	TargetTime
	TimerActive
	TimerSet
	Title
)
