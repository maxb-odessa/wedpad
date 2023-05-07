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
func (evHandler EventHandler) ReservoirReplenished(eventData map[string]interface{}) {
    // ev := new(evReservoirReplenished)
    // mapstructure.Decode(eventData, ev)
}

