package events

import (
	"time"
)

// Music event structure
type MusicT struct {
	MusicTrack string    `mapstructure:"MusicTrack"`
	Event      string    `mapstructure:"event"`
	Timestamp  time.Time `mapstructure:"timestamp"`
}

// Music event handler
func (evHandler EventHandler) Music(eventData map[string]interface{}) {
    // ev := new(MusicT)
    // mapstructure.Decode(eventData, ev)
}

