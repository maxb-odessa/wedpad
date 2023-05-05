package events

import (
	"time"
)

// ReservoirReplenished event structure
type evReservoirReplenished struct {
	FuelMain      float64   `mapstructure:"FuelMain"`
	FuelReservoir float64   `mapstructure:"FuelReservoir"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// ReservoirReplenished event handler
func ReservoirReplenished(e interface{}) {
    // ev := e.(evReservoirReplenished)
}

