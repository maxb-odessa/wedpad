package events

import (
	"time"
)

// Liftoff event structure
type evLiftoff struct {
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

// Liftoff event handler
func Liftoff(e interface{}) {
    // ev := e.(evLiftoff)
}

