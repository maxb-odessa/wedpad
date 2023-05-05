package events

import (
	"time"
)

// RedeemVoucher event structure
type evRedeemVoucher struct {
	Amount           int       `mapstructure:"Amount"`
	BrokerPercentage float64   `mapstructure:"BrokerPercentage"`
	Type             string    `mapstructure:"Type"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// RedeemVoucher event handler
func RedeemVoucher(e interface{}) {
    // ev := e.(evRedeemVoucher)
}

