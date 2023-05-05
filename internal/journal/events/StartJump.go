package events

import (
	"time"
)

// StartJump event structure
type evStartJump struct {
	JumpType      string    `mapstructure:"JumpType"`
	StarClass     string    `mapstructure:"StarClass"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// StartJump event handler
func StartJump(eventData map[string]interface{}) {
    // ev := new(evStartJump)
    // mapstructure.Decode(eventData, ev)
}

