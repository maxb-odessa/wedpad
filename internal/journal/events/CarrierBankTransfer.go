package events

import (
	"time"
)

// CarrierBankTransfer event structure
type evCarrierBankTransfer struct {
	CarrierBalance int       `mapstructure:"CarrierBalance"`
	CarrierID      int       `mapstructure:"CarrierID"`
	Deposit        int       `mapstructure:"Deposit"`
	PlayerBalance  int       `mapstructure:"PlayerBalance"`
	Withdraw       int       `mapstructure:"Withdraw"`
	Event          string    `mapstructure:"event"`
	Timestamp      time.Time `mapstructure:"timestamp"`
}

// CarrierBankTransfer event handler
func CarrierBankTransfer(e interface{}) {
    // ev := e.(evCarrierBankTransfer)
}

