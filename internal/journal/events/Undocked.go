package events

import (
	"time"
)

// Undocked event structure
type UndockedT struct {
	MarketID    int       `mapstructure:"MarketID"`
	Multicrew   bool      `mapstructure:"Multicrew"`
	StationName string    `mapstructure:"StationName"`
	StationType string    `mapstructure:"StationType"`
	Taxi        bool      `mapstructure:"Taxi"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// Undocked event handler
func (evHandler EventHandler) Undocked(eventData map[string]interface{}) {
    // ev := new(UndockedT)
    // mapstructure.Decode(eventData, ev)
}

