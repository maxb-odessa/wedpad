package events

import (
	"time"
)

// DataScanned event structure
type evDataScanned struct {
	Type          string    `mapstructure:"Type"`
	TypeLocalised string    `mapstructure:"Type_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// DataScanned event handler
func DataScanned(e interface{}) {
    // ev := e.(evDataScanned)
}

