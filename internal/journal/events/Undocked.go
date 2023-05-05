package events

import (
	"time"
)

// Undocked event structure
type evUndocked struct {
	MarketID    int       `mapstructure:"MarketID"`
	Multicrew   bool      `mapstructure:"Multicrew"`
	StationName string    `mapstructure:"StationName"`
	StationType string    `mapstructure:"StationType"`
	Taxi        bool      `mapstructure:"Taxi"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// Undocked event handler
func Undocked(e interface{}) {
    // ev := e.(evUndocked)
}

