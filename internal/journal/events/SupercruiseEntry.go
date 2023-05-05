package events

import (
	"time"
)

// SupercruiseEntry event structure
type evSupercruiseEntry struct {
	Multicrew     bool      `mapstructure:"Multicrew"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Taxi          bool      `mapstructure:"Taxi"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// SupercruiseEntry event handler
func SupercruiseEntry(e interface{}) {
    // ev := e.(evSupercruiseEntry)
}

