package events

import (
	"time"
)

// RestockVehicle event structure
type evRestockVehicle struct {
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
func RestockVehicle(eventData map[string]interface{}) {
    // ev := new(evRestockVehicle)
    // mapstructure.Decode(eventData, ev)
}

