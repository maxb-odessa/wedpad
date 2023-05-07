package events

import (
	"time"
)

// CarrierDepositFuel event structure
type CarrierDepositFuelT struct {
	Amount    int       `mapstructure:"Amount"`
	CarrierID int       `mapstructure:"CarrierID"`
	Total     int       `mapstructure:"Total"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CarrierDepositFuel event handler
func (evHandler EventHandler) CarrierDepositFuel(eventData map[string]interface{}) {
    // ev := new(CarrierDepositFuelT)
    // mapstructure.Decode(eventData, ev)
}

