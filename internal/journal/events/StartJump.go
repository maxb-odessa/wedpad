package events

import (
	"time"

	"github.com/mitchellh/mapstructure"
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
	ev := new(StartJumpT)
	mapstructure.Decode(eventData, ev)

	cs := CurrentSystem
	if ev.JumpType == "Hyperspace" {
		cs.Name = ev.StarSystem
		cs.MainStarType = ev.StarClass
	}
}
