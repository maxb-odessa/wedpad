package events

import (
	"time"
)

// DatalinkScan event structure
type DatalinkScanT struct {
	Message          string    `mapstructure:"Message"`
	MessageLocalised string    `mapstructure:"Message_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// DatalinkScan event handler
func (evHandler EventHandler) DatalinkScan(eventData map[string]interface{}) {
    // ev := new(DatalinkScanT)
    // mapstructure.Decode(eventData, ev)
}

