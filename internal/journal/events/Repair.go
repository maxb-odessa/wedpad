package events

import (
	"time"
)

// Repair event structure
type RepairT struct {
	Cost      int       `mapstructure:"Cost"`
	Items     []string  `mapstructure:"Items"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Repair event handler
func (evHandler EventHandler) Repair(eventData map[string]interface{}) {
    // ev := new(RepairT)
    // mapstructure.Decode(eventData, ev)
}

