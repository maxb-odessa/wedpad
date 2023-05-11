package events

import (
	"time"
)

// ModuleInfo event structure
type ModuleInfoT struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// ModuleInfo event handler
func (evh *EventHandler) ModuleInfo(eventData map[string]interface{}) {
    // ev := new(ModuleInfoT)
    // mapstructure.Decode(eventData, ev)
}

