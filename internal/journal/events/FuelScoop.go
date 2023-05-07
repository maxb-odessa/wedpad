package events

import (
	"time"
)

// FuelScoop event structure
type evFuelScoop struct {
	Scooped   float64   `mapstructure:"Scooped"`
	Total     float64   `mapstructure:"Total"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// FuelScoop event handler
func (evHandler EventHandler) FuelScoop(eventData map[string]interface{}) {
    // ev := new(evFuelScoop)
    // mapstructure.Decode(eventData, ev)
}

