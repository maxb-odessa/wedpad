package events

import (
	"time"
)

// LeaveBody event structure
type evLeaveBody struct {
	Body          string    `mapstructure:"Body"`
	BodyID        int       `mapstructure:"BodyID"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// LeaveBody event handler
func LeaveBody(e interface{}) {
    // ev := e.(evLeaveBody)
}

