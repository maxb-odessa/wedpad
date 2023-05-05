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
func CarrierJumpCancelled(e interface{}) {
    // ev := e.(evCarrierJumpCancelled)
}

