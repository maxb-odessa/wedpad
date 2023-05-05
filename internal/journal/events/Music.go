package events

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

// Music event structure
type evMusic struct {
	MusicTrack string    `mapstructure:"MusicTrack"`
	Event      string    `mapstructure:"event"`
	Timestamp  time.Time `mapstructure:"timestamp"`
}

// Music event handler
func Music(eventData map[string]interface{}) {
	ev := new(evMusic)
	mapstructure.Decode(eventData, ev)
}
