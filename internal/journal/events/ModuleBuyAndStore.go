package events

import (
	"time"
)

// ModuleBuyAndStore event structure
type evModuleBuyAndStore struct {
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
func ModuleBuyAndStore(e interface{}) {
    // ev := e.(evModuleBuyAndStore)
}

