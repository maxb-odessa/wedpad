package events

import (
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// StartJump event structure
type StartJumpT struct {
	JumpType      string    `mapstructure:"JumpType"`
	StarClass     string    `mapstructure:"StarClass"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Taxi          bool      `mapstructure:"Taxi"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// StartJump event handler
func (evh *EventHandler) StartJump(eventData map[string]interface{}) {
	ev := new(StartJumpT)
	mapstructure.Decode(eventData, ev)

	if ev.JumpType != "Hyperspace" {
		return
	}

	cs := evh.CurrentSystem()

	cs.Reset()

	cs.SetName(ev.StarSystem)
	cs.SetMainStarType(ev.StarClass)

	m := &msg.Message{
		Target: msg.TARGET_LOG,
		Action: msg.ACTION_APPEND,
		Type:   msg.TYPE_VIEW,
		Data:   "Jumping to system " + ev.StarSystem + ", star class " + ev.StarClass + "<br>",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_LOG,
		Action: msg.ACTION_ATTENTION,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}
	m.Send()
}
