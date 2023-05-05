package events

import (
	"time"
)

// Missions event structure
type evMissions struct {
	Active    []interface{} `mapstructure:"Active"`
	Complete  []interface{} `mapstructure:"Complete"`
	Failed    []interface{} `mapstructure:"Failed"`
	Event     string        `mapstructure:"event"`
	Timestamp time.Time     `mapstructure:"timestamp"`
}

// Missions event handler
func Missions(e interface{}) {
    // ev := e.(evMissions)
}

