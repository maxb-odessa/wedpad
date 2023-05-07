package events

import (
	"time"
)

// CarrierJumpCancelled event structure
type evCarrierJumpCancelled struct {
	CarrierID int       `mapstructure:"CarrierID"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CarrierJumpCancelled event handler
func (evHandler EventHandler) CarrierJumpCancelled(eventData map[string]interface{}) {
    // ev := new(evCarrierJumpCancelled)
    // mapstructure.Decode(eventData, ev)
}

