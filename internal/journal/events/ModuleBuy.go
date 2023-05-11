package events

import (
	"time"
)

// ModuleBuy event structure
type ModuleBuyT struct {
	BuyItem             string    `mapstructure:"BuyItem"`
	BuyItemLocalised    string    `mapstructure:"BuyItem_Localised"`
	BuyPrice            int       `mapstructure:"BuyPrice"`
	MarketID            int       `mapstructure:"MarketID"`
	Ship                string    `mapstructure:"Ship"`
	ShipID              int       `mapstructure:"ShipID"`
	Slot                string    `mapstructure:"Slot"`
	StoredItem          string    `mapstructure:"StoredItem"`
	StoredItemLocalised string    `mapstructure:"StoredItem_Localised"`
	Event               string    `mapstructure:"event"`
	Timestamp           time.Time `mapstructure:"timestamp"`
}

// ModuleBuy event handler
func (evh *EventHandler) ModuleBuy(eventData map[string]interface{}) {
    // ev := new(ModuleBuyT)
    // mapstructure.Decode(eventData, ev)
}

