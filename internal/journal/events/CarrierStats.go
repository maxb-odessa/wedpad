package events

import (
	"time"
)

// CarrierStats event structure
type CarrierStatsT struct {
	AllowNotorious bool   `mapstructure:"AllowNotorious"`
	Callsign       string `mapstructure:"Callsign"`
	CarrierID      int    `mapstructure:"CarrierID"`
	Crew           []struct {
		Activated bool   `mapstructure:"Activated"`
		CrewName  string `mapstructure:"CrewName"`
		CrewRole  string `mapstructure:"CrewRole"`
		Enabled   bool   `mapstructure:"Enabled"`
	} `mapstructure:"Crew"`
	DockingAccess string `mapstructure:"DockingAccess"`
	Finance       struct {
		AvailableBalance int `mapstructure:"AvailableBalance"`
		CarrierBalance   int `mapstructure:"CarrierBalance"`
		ReserveBalance   int `mapstructure:"ReserveBalance"`
		ReservePercent   int `mapstructure:"ReservePercent"`
		TaxRateRearm     int `mapstructure:"TaxRate_rearm"`
		TaxRateRefuel    int `mapstructure:"TaxRate_refuel"`
		TaxRateRepair    int `mapstructure:"TaxRate_repair"`
	} `mapstructure:"Finance"`
	FuelLevel           int           `mapstructure:"FuelLevel"`
	JumpRangeCurr       float64       `mapstructure:"JumpRangeCurr"`
	JumpRangeMax        float64       `mapstructure:"JumpRangeMax"`
	ModulePacks         []interface{} `mapstructure:"ModulePacks"`
	Name                string        `mapstructure:"Name"`
	PendingDecommission bool          `mapstructure:"PendingDecommission"`
	ShipPacks           []interface{} `mapstructure:"ShipPacks"`
	SpaceUsage          struct {
		Cargo              int `mapstructure:"Cargo"`
		CargoSpaceReserved int `mapstructure:"CargoSpaceReserved"`
		Crew               int `mapstructure:"Crew"`
		FreeSpace          int `mapstructure:"FreeSpace"`
		ModulePacks        int `mapstructure:"ModulePacks"`
		ShipPacks          int `mapstructure:"ShipPacks"`
		TotalCapacity      int `mapstructure:"TotalCapacity"`
	} `mapstructure:"SpaceUsage"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CarrierStats event handler
func (evh *EventHandler) CarrierStats(eventData map[string]interface{}) {
    // ev := new(CarrierStatsT)
    // mapstructure.Decode(eventData, ev)
}

