package events

import (
	"time"
)

// CarrierBuy event structure
type evCarrierBuy struct {
	BoughtAtMarket int       `mapstructure:"BoughtAtMarket"`
	Callsign       string    `mapstructure:"Callsign"`
	CarrierID      int       `mapstructure:"CarrierID"`
	Location       string    `mapstructure:"Location"`
	Price          int       `mapstructure:"Price"`
	SystemAddress  int       `mapstructure:"SystemAddress"`
	Variant        string    `mapstructure:"Variant"`
	Event          string    `mapstructure:"event"`
	Timestamp      time.Time `mapstructure:"timestamp"`
}

// CarrierBuy event handler
func (evHandler EventHandler) CarrierBuy(eventData map[string]interface{}) {
    // ev := new(evCarrierBuy)
    // mapstructure.Decode(eventData, ev)
}

