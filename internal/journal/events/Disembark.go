package events

import (
	"time"
)

// Disembark event structure
type DisembarkT struct {
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

// Disembark event handler
func (evHandler EventHandler) Disembark(eventData map[string]interface{}) {
    // ev := new(DisembarkT)
    // mapstructure.Decode(eventData, ev)
}

