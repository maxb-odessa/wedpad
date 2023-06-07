package events

import (
	"time"

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

	showNavRoute(cs.Name(), cs.MainStarType(), 0, "", "")
}
