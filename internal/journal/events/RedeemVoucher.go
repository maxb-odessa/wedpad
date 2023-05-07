package events

import (
	"time"
)

// RedeemVoucher event structure
type RedeemVoucherT struct {
	Amount           int       `mapstructure:"Amount"`
	BrokerPercentage float64   `mapstructure:"BrokerPercentage"`
	Type             string    `mapstructure:"Type"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// RedeemVoucher event handler
func (evHandler EventHandler) RedeemVoucher(eventData map[string]interface{}) {
    // ev := new(RedeemVoucherT)
    // mapstructure.Decode(eventData, ev)
}

