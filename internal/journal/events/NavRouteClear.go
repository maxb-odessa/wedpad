package events

import (
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// NavRouteClear event structure
type NavRouteClearT struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// NavRouteClear event handler
func (evh *EventHandler) NavRouteClear(eventData map[string]interface{}) {
	ev := new(NavRouteClearT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()

	type JumpRoute struct {
		PrevSystem, NextSystem string
		PrevStar, NextStar     TypeColorPair
		Jumps                  int
	}

	Data := &JumpRoute{
		PrevSystem: cs.Name(),
		PrevStar:   StarTypeColor(cs.MainStarType()),
	}

	m := &msg.Message{
		Action: msg.ACTION_REPLACE,
		Target: msg.TARGET_TOP,
		Type:   msg.TYPE_VIEW,
		Data:   Data,
	}

	m.Send()

}
