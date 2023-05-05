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
func Music(e interface{}) {
    // ev := e.(evMusic)
}

