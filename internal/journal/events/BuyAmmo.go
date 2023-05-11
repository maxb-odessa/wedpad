package events

import (
	"time"
)

// BuyAmmo event structure
type BuyAmmoT struct {
	Cost      int       `mapstructure:"Cost"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// BuyAmmo event handler
func (evh *EventHandler) BuyAmmo(eventData map[string]interface{}) {
    // ev := new(BuyAmmoT)
    // mapstructure.Decode(eventData, ev)
}

