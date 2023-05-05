package events

import (
	"time"
)

// ShipyardNew event structure
type evShipyardNew struct {
	NewShipID         int       `mapstructure:"NewShipID"`
	ShipType          string    `mapstructure:"ShipType"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ShipyardNew event handler
func ShipyardNew(eventData map[string]interface{}) {
    // ev := new(evShipyardNew)
    // mapstructure.Decode(eventData, ev)
}

