package events

import (
	"time"
)

// ShipyardBuy event structure
type ShipyardBuyT struct {
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
func (evh *EventHandler) ShipyardBuy(eventData map[string]interface{}) {
    // ev := new(ShipyardBuyT)
    // mapstructure.Decode(eventData, ev)
}

