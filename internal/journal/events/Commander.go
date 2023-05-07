package events

import (
	"time"
)

// Commander event structure
type CommanderT struct {
	Fid       string    `mapstructure:"FID"`
	Name      string    `mapstructure:"Name"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Commander event handler
func (evHandler EventHandler) Commander(eventData map[string]interface{}) {
    // ev := new(CommanderT)
    // mapstructure.Decode(eventData, ev)
}

