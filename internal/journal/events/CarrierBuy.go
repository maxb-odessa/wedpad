package events

import (
	"time"
)

// CarrierBuy event structure
type CarrierBuyT struct {
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
func (evh *EventHandler) CarrierBuy(eventData map[string]interface{}) {
    // ev := new(CarrierBuyT)
    // mapstructure.Decode(eventData, ev)
    // cs := evh.CurrentSystem()
}

