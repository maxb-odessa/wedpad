package events

import (
	"time"
)

// CollectCargo event structure
type evCollectCargo struct {
	Stolen        bool      `mapstructure:"Stolen"`
	Type          string    `mapstructure:"Type"`
	TypeLocalised string    `mapstructure:"Type_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// CollectCargo event handler
func (evHandler EventHandler) CollectCargo(eventData map[string]interface{}) {
    // ev := new(evCollectCargo)
    // mapstructure.Decode(eventData, ev)
}

