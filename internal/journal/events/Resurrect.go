package events

import (
	"time"
)

// Resurrect event structure
type ResurrectT struct {
	Bankrupt  bool      `mapstructure:"Bankrupt"`
	Cost      int       `mapstructure:"Cost"`
	Option    string    `mapstructure:"Option"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Resurrect event handler
func (evh *EventHandler) Resurrect(eventData map[string]interface{}) {
    // ev := new(ResurrectT)
    // mapstructure.Decode(eventData, ev)
}

