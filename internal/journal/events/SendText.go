package events

import (
	"time"
)

// SendText event structure
type evSendText struct {
	Message   string    `mapstructure:"Message"`
	Sent      bool      `mapstructure:"Sent"`
	To        string    `mapstructure:"To"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// SendText event handler
func SendText(e interface{}) {
    // ev := e.(evSendText)
}

