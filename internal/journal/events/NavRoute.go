package events

import (
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// NavRoute event structure
// { "StarSystem":"Mylaifai BV-W d2-2620", "SystemAddress":90032907931803, "StarPos":[-413.59375,-357.31250,18440.93750], "StarClass":"N" },
type NavRouteT struct {
	Route []struct {
		StarSystem    string     `mapstructure:"StarSystem"`
		SystemAddress uint64     `mapstructure:"SystemAddress"`
		StarPos       [3]float64 `mapstructure:"StarPos"`
		StarClass     string     `mapstructure:"StarClass"`
	} `mapstructure:"Route"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// NavRoute event handler
func (evh *EventHandler) NavRoute(eventData map[string]interface{}) {
	ev := new(NavRouteT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()

	if jumps := len(ev.Route); jumps > 0 {
		jumps := jumps - 1
		showNavRoute(cs.Name(), cs.MainStarType(), jumps, ev.Route[jumps].StarSystem, ev.Route[jumps].StarClass)
	}

}

func showNavRoute(pSys, pStar string, jumpsLeft int, nSys, nStar string) {

	type JumpRoute struct {
		PrevSystem, NextSystem string
		PrevStar, NextStar     TypeColorPair
		Jumps                  int
	}

	Data := &JumpRoute{
		PrevSystem: pSys,
		PrevStar:   StarTypeColor(pStar),
		Jumps:      jumpsLeft,
		NextSystem: nSys,
		NextStar:   StarTypeColor(nStar),
	}

	m := &msg.Message{
		Action: msg.ACTION_REPLACE,
		Target: msg.TARGET_TOP,
		Type:   msg.TYPE_VIEW,
		Data:   Data,
	}

	m.Send()
}
