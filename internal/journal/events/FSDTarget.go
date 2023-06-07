package events

import (
	"time"

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
func (evh *EventHandler) FSDTarget(eventData map[string]interface{}) {
	ev := new(FSDTargetT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()

	showNavRoute(cs.Name(), cs.MainStarType(), ev.RemainingJumpsInRoute, ev.Name, ev.StarClass)

	/*
		cs.Reset()
		cs.SetMainStarType(ev.StarClass)
		cs.SetName(ev.Name)
	*/
}
