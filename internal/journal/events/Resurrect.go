package events

import (
	"time"
)

// Resurrect event structure
type evResurrect struct {
	Bankrupt  bool      `mapstructure:"Bankrupt"`
	Cost      int       `mapstructure:"Cost"`
	Option    string    `mapstructure:"Option"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Resurrect event handler
func (evHandler EventHandler) Resurrect(eventData map[string]interface{}) {
    // ev := new(evResurrect)
    // mapstructure.Decode(eventData, ev)
}

