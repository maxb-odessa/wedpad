package events

import (
	"time"
)

// DiscoveryScan event structure
type evDiscoveryScan struct {
	Bodies        int       `mapstructure:"Bodies"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// DiscoveryScan event handler
func (evHandler EventHandler) DiscoveryScan(eventData map[string]interface{}) {
    // ev := new(evDiscoveryScan)
    // mapstructure.Decode(eventData, ev)
}

