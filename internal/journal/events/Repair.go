package events

import (
	"time"
)

// Repair event structure
type evRepair struct {
	Cost      int       `mapstructure:"Cost"`
	Items     []string  `mapstructure:"Items"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Repair event handler
func Repair(eventData map[string]interface{}) {
    // ev := new(evRepair)
    // mapstructure.Decode(eventData, ev)
}

