package events

import (
	"time"
)

// DiscoveryScan event structure
type DiscoveryScanT struct {
	Bodies        int       `mapstructure:"Bodies"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// DiscoveryScan event handler
func (evHandler EventHandler) DiscoveryScan(eventData map[string]interface{}) {
    // ev := new(DiscoveryScanT)
    // mapstructure.Decode(eventData, ev)
}

