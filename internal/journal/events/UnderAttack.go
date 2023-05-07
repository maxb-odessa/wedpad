package events

import (
	"time"
)

// UnderAttack event structure
type UnderAttackT struct {
	Target    string    `mapstructure:"Target"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// UnderAttack event handler
func (evHandler EventHandler) UnderAttack(eventData map[string]interface{}) {
    // ev := new(UnderAttackT)
    // mapstructure.Decode(eventData, ev)
}

