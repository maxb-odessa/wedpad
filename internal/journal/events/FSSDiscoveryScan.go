package events

import (
	"time"
)

// FSSDiscoveryScan event structure
type evFSSDiscoveryScan struct {
	BodyCount     int       `mapstructure:"BodyCount"`
	NonBodyCount  int       `mapstructure:"NonBodyCount"`
	Progress      float64   `mapstructure:"Progress"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	SystemName    string    `mapstructure:"SystemName"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// FSSDiscoveryScan event handler
func FSSDiscoveryScan(eventData map[string]interface{}) {
    // ev := new(evFSSDiscoveryScan)
    // mapstructure.Decode(eventData, ev)
}

