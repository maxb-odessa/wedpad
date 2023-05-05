package events

import (
	"time"
)

// BuyMicroResources event structure
type evBuyMicroResources struct {
	Category      string    `mapstructure:"Category"`
	Count         int       `mapstructure:"Count"`
	MarketID      int       `mapstructure:"MarketID"`
	Name          string    `mapstructure:"Name"`
	NameLocalised string    `mapstructure:"Name_Localised"`
	Price         int       `mapstructure:"Price"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// BuyMicroResources event handler
func BuyMicroResources(e interface{}) {
    // ev := e.(evBuyMicroResources)
}

