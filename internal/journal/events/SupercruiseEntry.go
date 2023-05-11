package events

import (
	"time"
)

// SupercruiseEntry event structure
type SupercruiseEntryT struct {
	Multicrew     bool      `mapstructure:"Multicrew"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Taxi          bool      `mapstructure:"Taxi"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// SupercruiseEntry event handler
func (evh *EventHandler) SupercruiseEntry(eventData map[string]interface{}) {
    // ev := new(SupercruiseEntryT)
    // mapstructure.Decode(eventData, ev)
}

