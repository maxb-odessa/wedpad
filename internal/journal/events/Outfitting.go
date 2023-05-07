package events

import (
	"time"
)

// Outfitting event structure
type evOutfitting struct {
	Horizons    bool          `mapstructure:"Horizons"`
	Items       []interface{} `mapstructure:"Items"`
	MarketID    int           `mapstructure:"MarketID"`
	StarSystem  string        `mapstructure:"StarSystem"`
	StationName string        `mapstructure:"StationName"`
	Event       string        `mapstructure:"event"`
	Timestamp   time.Time     `mapstructure:"timestamp"`
}

// Outfitting event handler
func (evHandler EventHandler) Outfitting(eventData map[string]interface{}) {
    // ev := new(evOutfitting)
    // mapstructure.Decode(eventData, ev)
}
