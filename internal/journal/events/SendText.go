package events

import (
	"time"
)

// SendText event structure
type SendTextT struct {
	Message   string    `mapstructure:"Message"`
	Sent      bool      `mapstructure:"Sent"`
	To        string    `mapstructure:"To"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// SendText event handler
func (evHandler EventHandler) SendText(eventData map[string]interface{}) {
    // ev := new(SendTextT)
    // mapstructure.Decode(eventData, ev)
}

