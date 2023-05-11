package events

import (
	"time"
)

// CarrierTradeOrder event structure
type CarrierTradeOrderT struct {
	BlackMarket   bool      `mapstructure:"BlackMarket"`
	CancelTrade   bool      `mapstructure:"CancelTrade"`
	CarrierID     int       `mapstructure:"CarrierID"`
	Commodity     string    `mapstructure:"Commodity"`
	Price         int       `mapstructure:"Price"`
	PurchaseOrder int       `mapstructure:"PurchaseOrder"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// CarrierTradeOrder event handler
func (evh *EventHandler) CarrierTradeOrder(eventData map[string]interface{}) {
    // ev := new(CarrierTradeOrderT)
    // mapstructure.Decode(eventData, ev)
}

