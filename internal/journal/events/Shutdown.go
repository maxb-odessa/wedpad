package events

import (
	"time"
)

// Shutdown event structure
type evShutdown struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Shutdown event handler
func Shutdown(eventData map[string]interface{}) {
    // ev := new(evShutdown)
    // mapstructure.Decode(eventData, ev)
}

