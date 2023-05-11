package events

import (
	"time"
)

// Missions event structure
type MissionsT struct {
	Active    []interface{} `mapstructure:"Active"`
	Complete  []interface{} `mapstructure:"Complete"`
	Failed    []interface{} `mapstructure:"Failed"`
	Event     string        `mapstructure:"event"`
	Timestamp time.Time     `mapstructure:"timestamp"`
}

// Missions event handler
func (evh *EventHandler) Missions(eventData map[string]interface{}) {
    // ev := new(MissionsT)
    // mapstructure.Decode(eventData, ev)
}

