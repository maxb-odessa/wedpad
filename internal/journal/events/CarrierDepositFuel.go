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
func CarrierDepositFuel(e interface{}) {
    // ev := e.(evCarrierDepositFuel)
}

