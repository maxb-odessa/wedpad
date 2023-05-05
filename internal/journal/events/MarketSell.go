package events

import (
	"time"
)

// MarketSell event structure
type evMarketSell struct {
	AvgPricePaid int       `mapstructure:"AvgPricePaid"`
	Count        int       `mapstructure:"Count"`
	MarketID     int       `mapstructure:"MarketID"`
	SellPrice    int       `mapstructure:"SellPrice"`
	TotalSale    int       `mapstructure:"TotalSale"`
	Type         string    `mapstructure:"Type"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// MarketSell event handler
func MarketSell(e interface{}) {
    // ev := e.(evMarketSell)
}

