package events

import (
	"time"
)

// Cargo event structure
type evCargo struct {
	Count     int           `mapstructure:"Count"`
	Inventory []interface{} `mapstructure:"Inventory"`
	Vessel    string        `mapstructure:"Vessel"`
	Event     string        `mapstructure:"event"`
	Timestamp time.Time     `mapstructure:"timestamp"`
}

// Cargo event handler
func Cargo(eventData map[string]interface{}) {
    // ev := new(evCargo)
    // mapstructure.Decode(eventData, ev)
}

