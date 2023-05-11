package events

import (
	"time"
)

// HeatWarning event structure
type HeatWarningT struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// HeatWarning event handler
func (evh *EventHandler) HeatWarning(eventData map[string]interface{}) {
    // ev := new(HeatWarningT)
    // mapstructure.Decode(eventData, ev)
}

