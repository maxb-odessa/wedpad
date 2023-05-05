package events

import (
	"time"
)

// HullDamage event structure
type evHullDamage struct {
	Fighter     bool      `mapstructure:"Fighter"`
	Health      float64   `mapstructure:"Health"`
	PlayerPilot bool      `mapstructure:"PlayerPilot"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// HullDamage event handler
func HullDamage(e interface{}) {
    // ev := e.(evHullDamage)
}

