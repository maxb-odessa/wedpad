package events

import (
	"time"
)

// ModuleBuyAndStore event structure
type ModuleBuyAndStoreT struct {
	BuyItem          string    `mapstructure:"BuyItem"`
	BuyItemLocalised string    `mapstructure:"BuyItem_Localised"`
	BuyPrice         int       `mapstructure:"BuyPrice"`
	MarketID         int       `mapstructure:"MarketID"`
	Ship             string    `mapstructure:"Ship"`
	ShipID           int       `mapstructure:"ShipID"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// ModuleBuyAndStore event handler
func (evh *EventHandler) ModuleBuyAndStore(eventData map[string]interface{}) {
    // ev := new(ModuleBuyAndStoreT)
    // mapstructure.Decode(eventData, ev)
}

