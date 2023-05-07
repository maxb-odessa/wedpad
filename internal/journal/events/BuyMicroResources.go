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
func (evHandler EventHandler) BuyMicroResources(eventData map[string]interface{}) {
    // ev := new(evBuyMicroResources)
    // mapstructure.Decode(eventData, ev)
}

