package events

import (
	"time"
)

// BuyAmmo event structure
type evBuyAmmo struct {
	Cost      int       `mapstructure:"Cost"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// BuyAmmo event handler
func BuyAmmo(e interface{}) {
    // ev := e.(evBuyAmmo)
}

