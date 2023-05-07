package events

import (
	"time"
)

// DataScanned event structure
type DataScannedT struct {
	Type          string    `mapstructure:"Type"`
	TypeLocalised string    `mapstructure:"Type_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// DataScanned event handler
func (evHandler EventHandler) DataScanned(eventData map[string]interface{}) {
    // ev := new(DataScannedT)
    // mapstructure.Decode(eventData, ev)
}

