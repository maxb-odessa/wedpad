package events

import (
	"time"
)

// BuyMicroResources event structure
type BuyMicroResourcesT struct {
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
func (evh *EventHandler) BuyMicroResources(eventData map[string]interface{}) {
    // ev := new(BuyMicroResourcesT)
    // mapstructure.Decode(eventData, ev)
}

