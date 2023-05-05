package events

import (
	"time"
)

// HeatWarning event structure
type evHeatWarning struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// HeatWarning event handler
func HeatWarning(e interface{}) {
    // ev := e.(evHeatWarning)
}

