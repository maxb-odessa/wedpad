package events

import (
	"time"
)

// Fileheader event structure
type FileheaderT struct {
	Odyssey     bool      `mapstructure:"Odyssey"`
	Build       string    `mapstructure:"build"`
	Event       string    `mapstructure:"event"`
	Gameversion string    `mapstructure:"gameversion"`
	Language    string    `mapstructure:"language"`
	Part        int       `mapstructure:"part"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// Fileheader event handler
func (evHandler EventHandler) Fileheader(eventData map[string]interface{}) {
    // ev := new(FileheaderT)
    // mapstructure.Decode(eventData, ev)
}

