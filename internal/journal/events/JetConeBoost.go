package events

import (
	"time"
)

// JetConeBoost event structure
type evJetConeBoost struct {
	BoostValue float64   `mapstructure:"BoostValue"`
	Event      string    `mapstructure:"event"`
	Timestamp  time.Time `mapstructure:"timestamp"`
}

// JetConeBoost event handler
func JetConeBoost(e interface{}) {
    // ev := e.(evJetConeBoost)
}

