package events

import (
	"time"
)

// DockingDenied event structure
type evDockingDenied struct {
	MarketID    int       `mapstructure:"MarketID"`
	Reason      string    `mapstructure:"Reason"`
	StationName string    `mapstructure:"StationName"`
	StationType string    `mapstructure:"StationType"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// DockingDenied event handler
func (evHandler EventHandler) DockingDenied(eventData map[string]interface{}) {
    // ev := new(evDockingDenied)
    // mapstructure.Decode(eventData, ev)
}
