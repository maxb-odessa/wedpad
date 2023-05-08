package events

import (
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// FSDTarget event structure
type FSDTargetT struct {
	Name                  string    `mapstructure:"Name"`
	RemainingJumpsInRoute int       `mapstructure:"RemainingJumpsInRoute"`
	StarClass             string    `mapstructure:"StarClass"`
	SystemAddress         int       `mapstructure:"SystemAddress"`
	Event                 string    `mapstructure:"event"`
	Timestamp             time.Time `mapstructure:"timestamp"`
}

// FSDTarget event handler
func (evHandler EventHandler) FSDTarget(eventData map[string]interface{}) {
	ev := new(FSDTargetT)
	mapstructure.Decode(eventData, ev)

	cs := CurrentSystem

	type JumpRoute struct {
		PrevSystem, NextSystem string
		PrevStar, NextStar     StarTypeColorPair
		Jumps                  int
	}

	Data := &JumpRoute{
		PrevSystem: cs.Name,
		PrevStar:   StarTypeColor(cs.MainStarType()),
		Jumps:      ev.RemainingJumpsInRoute,
		NextSystem: ev.Name,
		NextStar:   StarTypeColor(ev.StarClass),
	}

	m := &msg.Message{
		Action: msg.ACTION_REPLACE,
		Target: msg.TARGET_TOP,
		Type:   msg.TYPE_VIEW,
		Data:   Data,
	}

	m.Send()
}
