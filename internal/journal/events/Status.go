package events

import (
	"time"
)

// Status event structure
type evStatus struct {
	Flags     int       `mapstructure:"Flags"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Status event handler
func Status(e interface{}) {
    // ev := e.(evStatus)
}

