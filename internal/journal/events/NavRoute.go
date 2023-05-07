package events

import (
	"time"
)

// NavRoute event structure
type evNavRoute struct {
	Route     []interface{} `mapstructure:"Route"`
	Event     string        `mapstructure:"event"`
	Timestamp time.Time     `mapstructure:"timestamp"`
}

// NavRoute event handler
func (evHandler EventHandler) NavRoute(eventData map[string]interface{}) {
    // ev := new(evNavRoute)
    // mapstructure.Decode(eventData, ev)
}

