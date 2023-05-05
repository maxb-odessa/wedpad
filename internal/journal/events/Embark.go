package events

import (
	"time"
)

// Embark event structure
type evEmbark struct {
	Body          string    `mapstructure:"Body"`
	BodyID        int       `mapstructure:"BodyID"`
	ID            int       `mapstructure:"ID"`
	MarketID      int       `mapstructure:"MarketID"`
	Multicrew     bool      `mapstructure:"Multicrew"`
	OnPlanet      bool      `mapstructure:"OnPlanet"`
	OnStation     bool      `mapstructure:"OnStation"`
	Srv           bool      `mapstructure:"SRV"`
	StarSystem    string    `mapstructure:"StarSystem"`
	StationName   string    `mapstructure:"StationName"`
	StationType   string    `mapstructure:"StationType"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Taxi          bool      `mapstructure:"Taxi"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// Embark event handler
func Embark(e interface{}) {
    // ev := e.(evEmbark)
}

