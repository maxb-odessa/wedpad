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
func DiscoveryScan(e interface{}) {
    // ev := e.(evDiscoveryScan)
}

