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
func (evHandler EventHandler) BuyAmmo(eventData map[string]interface{}) {
    // ev := new(evBuyAmmo)
    // mapstructure.Decode(eventData, ev)
}

