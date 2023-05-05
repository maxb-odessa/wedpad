package events

import (
	"time"
)

// FSDTarget event structure
type evFSDTarget struct {
	Name                  string    `mapstructure:"Name"`
	RemainingJumpsInRoute int       `mapstructure:"RemainingJumpsInRoute"`
	StarClass             string    `mapstructure:"StarClass"`
	SystemAddress         int       `mapstructure:"SystemAddress"`
	Event                 string    `mapstructure:"event"`
	Timestamp             time.Time `mapstructure:"timestamp"`
}

// FSDTarget event handler
func FSDTarget(eventData map[string]interface{}) {
    // ev := new(evFSDTarget)
    // mapstructure.Decode(eventData, ev)
}

