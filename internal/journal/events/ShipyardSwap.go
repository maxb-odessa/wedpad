package events

import (
	"time"
)

// ShipyardSwap event structure
type evShipyardSwap struct {
	MarketID          int       `mapstructure:"MarketID"`
	ShipID            int       `mapstructure:"ShipID"`
	ShipType          string    `mapstructure:"ShipType"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised"`
	StoreOldShip      string    `mapstructure:"StoreOldShip"`
	StoreShipID       int       `mapstructure:"StoreShipID"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ShipyardSwap event handler
func (evHandler EventHandler) ShipyardSwap(eventData map[string]interface{}) {
    // ev := new(evShipyardSwap)
    // mapstructure.Decode(eventData, ev)
}

