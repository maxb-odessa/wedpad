package events

import (
	"time"
)

// DockingGranted event structure
type evDockingGranted struct {
	LandingPad  int       `mapstructure:"LandingPad"`
	MarketID    int       `mapstructure:"MarketID"`
	StationName string    `mapstructure:"StationName"`
	StationType string    `mapstructure:"StationType"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// DockingGranted event handler
func DockingGranted(eventData map[string]interface{}) {
    // ev := new(evDockingGranted)
    // mapstructure.Decode(eventData, ev)
}

