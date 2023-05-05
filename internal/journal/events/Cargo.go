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
func Cargo(e interface{}) {
    // ev := e.(evCargo)
}

