package events

import (
	"time"
)

// Music event structure
type evMusic struct {
	MusicTrack string    `mapstructure:"MusicTrack"`
	Event      string    `mapstructure:"event"`
	Timestamp  time.Time `mapstructure:"timestamp"`
}

// Music event handler
func (evHandler EventHandler) Music(eventData map[string]interface{}) {
    // ev := new(evMusic)
    // mapstructure.Decode(eventData, ev)
}
