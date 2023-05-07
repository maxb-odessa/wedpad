package events

import (
	"time"
)

// ShipyardSell event structure
type evShipyardSell struct {
	MarketID          int       `mapstructure:"MarketID"`
	SellShipID        int       `mapstructure:"SellShipID"`
	ShipPrice         int       `mapstructure:"ShipPrice"`
	ShipType          string    `mapstructure:"ShipType"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ShipyardSell event handler
func (evHandler EventHandler) ShipyardSell(eventData map[string]interface{}) {
    // ev := new(evShipyardSell)
    // mapstructure.Decode(eventData, ev)
}

