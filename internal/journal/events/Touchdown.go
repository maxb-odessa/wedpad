package events

import (
	"time"
)

// Touchdown event structure
type TouchdownT struct {
	Body               string    `mapstructure:"Body"`
	BodyID             int       `mapstructure:"BodyID"`
	Latitude           float64   `mapstructure:"Latitude"`
	Longitude          float64   `mapstructure:"Longitude"`
	Multicrew          bool      `mapstructure:"Multicrew"`
	NearestDestination string    `mapstructure:"NearestDestination"`
	OnPlanet           bool      `mapstructure:"OnPlanet"`
	OnStation          bool      `mapstructure:"OnStation"`
	PlayerControlled   bool      `mapstructure:"PlayerControlled"`
	StarSystem         string    `mapstructure:"StarSystem"`
	SystemAddress      int       `mapstructure:"SystemAddress"`
	Taxi               bool      `mapstructure:"Taxi"`
	Event              string    `mapstructure:"event"`
	Timestamp          time.Time `mapstructure:"timestamp"`
}

// Touchdown event handler
func (evh *EventHandler) Touchdown(eventData map[string]interface{}) {
    // ev := new(TouchdownT)
    // mapstructure.Decode(eventData, ev)
}

