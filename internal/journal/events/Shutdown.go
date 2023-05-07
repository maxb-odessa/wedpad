package events

import (
	"time"
)

// Shutdown event structure
type ShutdownT struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Shutdown event handler
func (evHandler EventHandler) Shutdown(eventData map[string]interface{}) {
    // ev := new(ShutdownT)
    // mapstructure.Decode(eventData, ev)
}

