package events

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

// ReservoirReplenished event structure
type ReservoirReplenishedT struct {
	FuelMain      float64   `mapstructure:"FuelMain"`
	FuelReservoir float64   `mapstructure:"FuelReservoir"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// ReservoirReplenished event handler
func (evh *EventHandler) ReservoirReplenished(eventData map[string]interface{}) {
	ev := new(ReservoirReplenishedT)
	mapstructure.Decode(eventData, ev)

	AlertFuel(ev.FuelMain)
}
