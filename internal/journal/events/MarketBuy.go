package events

import (
	"time"
)

// MarketBuy event structure
type MarketBuyT struct {
	BuyPrice  int       `mapstructure:"BuyPrice"`
	Count     int       `mapstructure:"Count"`
	MarketID  int       `mapstructure:"MarketID"`
	TotalCost int       `mapstructure:"TotalCost"`
	Type      string    `mapstructure:"Type"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// MarketBuy event handler
func (evHandler EventHandler) MarketBuy(eventData map[string]interface{}) {
    // ev := new(MarketBuyT)
    // mapstructure.Decode(eventData, ev)
}

