package events

import (
	"time"
)

// ShipyardTransfer event structure
type ShipyardTransferT struct {
	Distance          float64   `mapstructure:"Distance"`
	MarketID          int       `mapstructure:"MarketID"`
	ShipID            int       `mapstructure:"ShipID"`
	ShipMarketID      int       `mapstructure:"ShipMarketID"`
	ShipType          string    `mapstructure:"ShipType"`
	ShipTypeLocalised string    `mapstructure:"ShipType_Localised"`
	System            string    `mapstructure:"System"`
	TransferPrice     int       `mapstructure:"TransferPrice"`
	TransferTime      int       `mapstructure:"TransferTime"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ShipyardTransfer event handler
func (evh *EventHandler) ShipyardTransfer(eventData map[string]interface{}) {
    // ev := new(ShipyardTransferT)
    // mapstructure.Decode(eventData, ev)
}

