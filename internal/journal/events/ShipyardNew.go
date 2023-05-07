package events

import (
	"time"
)

// ShipyardNew event structure
type ShipyardNewT struct {
	NewShipID         int       `mapstructure:"NewShipID"`
	ShipType          string    `mapstructure:"ShipType"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ShipyardNew event handler
func (evHandler EventHandler) ShipyardNew(eventData map[string]interface{}) {
    // ev := new(ShipyardNewT)
    // mapstructure.Decode(eventData, ev)
}

