package events

import (
	"time"
)

// CarrierJumpCancelled event structure
type CarrierJumpCancelledT struct {
	CarrierID int       `mapstructure:"CarrierID"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CarrierJumpCancelled event handler
func (evh *EventHandler) CarrierJumpCancelled(eventData map[string]interface{}) {
    // ev := new(CarrierJumpCancelledT)
    // mapstructure.Decode(eventData, ev)
    // cs := evh.CurrentSystem()
}

