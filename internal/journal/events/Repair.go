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
func Repair(e interface{}) {
    // ev := e.(evRepair)
}

