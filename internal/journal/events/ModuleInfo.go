package events

import (
	"time"
)

// ModuleInfo event structure
type evModuleInfo struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// ModuleInfo event handler
func (evHandler EventHandler) ModuleInfo(eventData map[string]interface{}) {
    // ev := new(evModuleInfo)
    // mapstructure.Decode(eventData, ev)
}

