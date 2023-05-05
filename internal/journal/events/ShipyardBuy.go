package events

import (
	"time"
)

// ShipyardBuy event structure
type evShipyardBuy struct {
	MarketID          int       `mapstructure:"MarketID"`
	ShipPrice         int       `mapstructure:"ShipPrice"`
	ShipType          string    `mapstructure:"ShipType"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised"`
	StoreOldShip      string    `mapstructure:"StoreOldShip"`
	StoreShipID       int       `mapstructure:"StoreShipID"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ShipyardBuy event handler
func ShipyardBuy(eventData map[string]interface{}) {
    // ev := new(evShipyardBuy)
    // mapstructure.Decode(eventData, ev)
}

