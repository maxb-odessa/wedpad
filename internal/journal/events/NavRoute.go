package events

import (
	"time"
)

// NavRoute event structure
type evNavRoute struct {
	Route []struct {
		StarClass     string    `mapstructure:"StarClass"`
		StarPos       []float64 `mapstructure:"StarPos"`
		StarSystem    string    `mapstructure:"StarSystem"`
		SystemAddress int       `mapstructure:"SystemAddress"`
	} `mapstructure:"Route"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// NavRoute event handler
func NavRoute(eventData map[string]interface{}) {
    // ev := new(evNavRoute)
    // mapstructure.Decode(eventData, ev)
}

