package events

import (
	"time"
)

// RestockVehicle event structure
type RestockVehicleT struct {
	Cost          int       `mapstructure:"Cost"`
	Count         int       `mapstructure:"Count"`
	ID            int       `mapstructure:"ID"`
	Loadout       string    `mapstructure:"Loadout"`
	Type          string    `mapstructure:"Type"`
	TypeLocalised string    `mapstructure:"Type_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// RestockVehicle event handler
func (evh *EventHandler) RestockVehicle(eventData map[string]interface{}) {
    // ev := new(RestockVehicleT)
    // mapstructure.Decode(eventData, ev)
}

