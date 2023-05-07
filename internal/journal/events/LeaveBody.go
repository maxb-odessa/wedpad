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
func (evHandler EventHandler) LeaveBody(eventData map[string]interface{}) {
    // ev := new(evLeaveBody)
    // mapstructure.Decode(eventData, ev)
}

