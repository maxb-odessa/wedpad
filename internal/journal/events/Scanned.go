package events

import (
	"time"
)

// Scanned event structure
type evScanned struct {
	ScanType  string    `mapstructure:"ScanType"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Scanned event handler
func (evHandler EventHandler) Scanned(eventData map[string]interface{}) {
    // ev := new(evScanned)
    // mapstructure.Decode(eventData, ev)
}

