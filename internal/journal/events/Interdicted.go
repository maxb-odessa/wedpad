package events

import (
	"time"
)

// Interdicted event structure
type InterdictedT struct {
	Faction     string    `mapstructure:"Faction"`
	Interdictor string    `mapstructure:"Interdictor"`
	IsPlayer    bool      `mapstructure:"IsPlayer"`
	Submitted   bool      `mapstructure:"Submitted"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// Interdicted event handler
func (evh *EventHandler) Interdicted(eventData map[string]interface{}) {
    // ev := new(InterdictedT)
    // mapstructure.Decode(eventData, ev)
}

