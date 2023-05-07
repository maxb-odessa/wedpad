package events

import (
	"time"
)

// MissionRedirected event structure
type MissionRedirectedT struct {
	MissionID             int       `mapstructure:"MissionID"`
	Name                  string    `mapstructure:"Name"`
	NewDestinationStation string    `mapstructure:"NewDestinationStation"`
	NewDestinationSystem  string    `mapstructure:"NewDestinationSystem"`
	OldDestinationStation string    `mapstructure:"OldDestinationStation"`
	OldDestinationSystem  string    `mapstructure:"OldDestinationSystem"`
	Event                 string    `mapstructure:"event"`
	Timestamp             time.Time `mapstructure:"timestamp"`
}

// MissionRedirected event handler
func (evHandler EventHandler) MissionRedirected(eventData map[string]interface{}) {
    // ev := new(MissionRedirectedT)
    // mapstructure.Decode(eventData, ev)
}

