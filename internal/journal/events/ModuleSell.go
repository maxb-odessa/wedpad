package events

import (
	"time"
)

// ModuleSell event structure
type evModuleSell struct {
	MarketID          int       `mapstructure:"MarketID"`
	SellItem          string    `mapstructure:"SellItem"`
	SellItemLocalised string    `mapstructure:"SellItem_Localised"`
	SellPrice         int       `mapstructure:"SellPrice"`
	Ship              string    `mapstructure:"Ship"`
	ShipID            int       `mapstructure:"ShipID"`
	Slot              string    `mapstructure:"Slot"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ModuleSell event handler
func ModuleSell(eventData map[string]interface{}) {
    // ev := new(evModuleSell)
    // mapstructure.Decode(eventData, ev)
}

