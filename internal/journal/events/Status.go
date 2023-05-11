package events

import (
	"time"
)

// Status event structure
type StatusT struct {
	Flags     int       `mapstructure:"Flags"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Status event handler
func (evh *EventHandler) Status(eventData map[string]interface{}) {
    // ev := new(StatusT)
    // mapstructure.Decode(eventData, ev)
}

