package events

import (
	"time"
)

// SupercruiseExit event structure
type evSupercruiseExit struct {
	Body          string    `mapstructure:"Body"`
	BodyID        int       `mapstructure:"BodyID"`
	BodyType      string    `mapstructure:"BodyType"`
	Multicrew     bool      `mapstructure:"Multicrew"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Taxi          bool      `mapstructure:"Taxi"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// SupercruiseExit event handler
func SupercruiseExit(e interface{}) {
    // ev := e.(evSupercruiseExit)
}

