package events

import (
	"time"
)

// FSSAllBodiesFound event structure
type FSSAllBodiesFoundT struct {
	Count         int       `mapstructure:"Count"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	SystemName    string    `mapstructure:"SystemName"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// FSSAllBodiesFound event handler
func (evh *EventHandler) FSSAllBodiesFound(eventData map[string]interface{}) {
	// ev := new(FSSAllBodiesFoundT)
	// mapstructure.Decode(eventData, ev)

	evh.cs.ShowPlanets()
	evh.cs.ShowSignals()
	evh.cs.ShowNotes()
}
