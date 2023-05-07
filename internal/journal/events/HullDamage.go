package events

import (
	"time"
)

// HullDamage event structure
type HullDamageT struct {
	Fighter     bool      `mapstructure:"Fighter"`
	Health      float64   `mapstructure:"Health"`
	PlayerPilot bool      `mapstructure:"PlayerPilot"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// HullDamage event handler
func (evHandler EventHandler) HullDamage(eventData map[string]interface{}) {
    // ev := new(HullDamageT)
    // mapstructure.Decode(eventData, ev)
}

