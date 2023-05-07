package events

import (
	"time"
)

// StartJump event structure
type StartJumpT struct {
	JumpType      string    `mapstructure:"JumpType"`
	StarClass     string    `mapstructure:"StarClass"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// StartJump event handler
func (evHandler EventHandler) StartJump(eventData map[string]interface{}) {
    // ev := new(StartJumpT)
    // mapstructure.Decode(eventData, ev)
}

