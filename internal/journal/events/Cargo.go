package events

import (
	"time"
)

// Cargo event structure
type CargoT struct {
	Count     int           `mapstructure:"Count"`
	Inventory []interface{} `mapstructure:"Inventory"`
	Vessel    string        `mapstructure:"Vessel"`
	Event     string        `mapstructure:"event"`
	Timestamp time.Time     `mapstructure:"timestamp"`
}

// Cargo event handler
func (evh *EventHandler) Cargo(eventData map[string]interface{}) {
    // ev := new(CargoT)
    // mapstructure.Decode(eventData, ev)
}

