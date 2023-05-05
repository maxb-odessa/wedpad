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
func Scanned(e interface{}) {
    // ev := e.(evScanned)
}

