package events

import (
	"time"
)

// ShieldState event structure
type ShieldStateT struct {
	ShieldsUp bool      `mapstructure:"ShieldsUp"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// ShieldState event handler
func (evh *EventHandler) ShieldState(eventData map[string]interface{}) {
    // ev := new(ShieldStateT)
    // mapstructure.Decode(eventData, ev)
}

