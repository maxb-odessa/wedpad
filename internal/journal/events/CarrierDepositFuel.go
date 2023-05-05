package events

import (
	"time"
)

// CarrierDepositFuel event structure
type evCarrierDepositFuel struct {
	Amount    int       `mapstructure:"Amount"`
	CarrierID int       `mapstructure:"CarrierID"`
	Total     int       `mapstructure:"Total"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CarrierDepositFuel event handler
func CarrierDepositFuel(eventData map[string]interface{}) {
    // ev := new(evCarrierDepositFuel)
    // mapstructure.Decode(eventData, ev)
}

