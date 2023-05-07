package events

import (
	"time"
)

// Scanned event structure
type ScannedT struct {
	ScanType  string    `mapstructure:"ScanType"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Scanned event handler
func (evHandler EventHandler) Scanned(eventData map[string]interface{}) {
    // ev := new(ScannedT)
    // mapstructure.Decode(eventData, ev)
}

