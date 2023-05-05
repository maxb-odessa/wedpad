package events

import (
	"time"
)

// ShieldState event structure
type evShieldState struct {
	ShieldsUp bool      `mapstructure:"ShieldsUp"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// ShieldState event handler
func ShieldState(e interface{}) {
    // ev := e.(evShieldState)
}

