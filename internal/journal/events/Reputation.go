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
func Reputation(eventData map[string]interface{}) {
    // ev := new(evReputation)
    // mapstructure.Decode(eventData, ev)
}

