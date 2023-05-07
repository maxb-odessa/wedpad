package events

import (
	"time"
)

// NavRouteClear event structure
type NavRouteClearT struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// NavRouteClear event handler
func (evHandler EventHandler) NavRouteClear(eventData map[string]interface{}) {
    // ev := new(NavRouteClearT)
    // mapstructure.Decode(eventData, ev)
}

