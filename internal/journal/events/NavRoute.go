package events

import (
	"time"
)

// NavRoute event structure
type NavRouteT struct {
	Route     []interface{} `mapstructure:"Route"`
	Event     string        `mapstructure:"event"`
	Timestamp time.Time     `mapstructure:"timestamp"`
}

// NavRoute event handler
func (evh *EventHandler) NavRoute(eventData map[string]interface{}) {
    // ev := new(NavRouteT)
    // mapstructure.Decode(eventData, ev)
}

