package events

import (
	"time"
)

// Commander event structure
type evCommander struct {
	Fid       string    `mapstructure:"FID"`
	Name      string    `mapstructure:"Name"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Commander event handler
func Commander(e interface{}) {
    // ev := e.(evCommander)
}

