package events

import (
	"time"
)

// NavRouteClear event structure
type evNavRouteClear struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// NavRouteClear event handler
func NavRouteClear(e interface{}) {
    // ev := e.(evNavRouteClear)
}

