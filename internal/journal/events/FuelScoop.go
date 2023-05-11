package events

import (
	"time"
)

// FuelScoop event structure
type FuelScoopT struct {
	Scooped   float64   `mapstructure:"Scooped"`
	Total     float64   `mapstructure:"Total"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// FuelScoop event handler
func (evh *EventHandler) FuelScoop(eventData map[string]interface{}) {
    // ev := new(FuelScoopT)
    // mapstructure.Decode(eventData, ev)
}

