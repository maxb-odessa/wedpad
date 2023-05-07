package events

import (
	"time"
)

// Shipyard event structure
type ShipyardT struct {
	AllowCobraMkIv bool          `mapstructure:"AllowCobraMkIV"`
	Horizons       bool          `mapstructure:"Horizons"`
	MarketID       int           `mapstructure:"MarketID"`
	PriceList      []interface{} `mapstructure:"PriceList"`
	StarSystem     string        `mapstructure:"StarSystem"`
	StationName    string        `mapstructure:"StationName"`
	Event          string        `mapstructure:"event"`
	Timestamp      time.Time     `mapstructure:"timestamp"`
}

// Shipyard event handler
func (evHandler EventHandler) Shipyard(eventData map[string]interface{}) {
    // ev := new(ShipyardT)
    // mapstructure.Decode(eventData, ev)
}

