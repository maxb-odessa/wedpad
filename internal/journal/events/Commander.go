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
func Commander(eventData map[string]interface{}) {
    // ev := new(evCommander)
    // mapstructure.Decode(eventData, ev)
}

