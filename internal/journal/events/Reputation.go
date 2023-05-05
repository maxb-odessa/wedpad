package events

import (
	"time"
)

// Reputation event structure
type evReputation struct {
	Alliance   float64   `mapstructure:"Alliance"`
	Empire     float64   `mapstructure:"Empire"`
	Federation float64   `mapstructure:"Federation"`
	Event      string    `mapstructure:"event"`
	Timestamp  time.Time `mapstructure:"timestamp"`
}

// Reputation event handler
func Reputation(e interface{}) {
    // ev := e.(evReputation)
}

