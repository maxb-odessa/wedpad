package events

import (
	"time"
)

// FSSDiscoveryScan event structure
type FSSDiscoveryScanT struct {
	BodyCount     int       `mapstructure:"BodyCount"`
	NonBodyCount  int       `mapstructure:"NonBodyCount"`
	Progress      float64   `mapstructure:"Progress"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	SystemName    string    `mapstructure:"SystemName"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// FSSDiscoveryScan event handler
func (evHandler EventHandler) FSSDiscoveryScan(eventData map[string]interface{}) {
    // ev := new(FSSDiscoveryScanT)
    // mapstructure.Decode(eventData, ev)
}

