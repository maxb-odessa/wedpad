package events

import (
	"time"
)

// DatalinkScan event structure
type evDatalinkScan struct {
	Message          string    `mapstructure:"Message"`
	MessageLocalised string    `mapstructure:"Message_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// DatalinkScan event handler
func DatalinkScan(e interface{}) {
    // ev := e.(evDatalinkScan)
}

