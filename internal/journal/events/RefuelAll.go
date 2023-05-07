package events

import (
	"time"
)

// RefuelAll event structure
type evRefuelAll struct {
	Amount    float64   `mapstructure:"Amount"`
	Cost      int       `mapstructure:"Cost"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// RefuelAll event handler
func (evHandler EventHandler) RefuelAll(eventData map[string]interface{}) {
    // ev := new(evRefuelAll)
    // mapstructure.Decode(eventData, ev)
}

