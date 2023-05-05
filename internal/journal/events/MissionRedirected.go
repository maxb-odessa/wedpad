package events

import (
	"time"
)

// MissionRedirected event structure
type evMissionRedirected struct {
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
func MissionRedirected(e interface{}) {
    // ev := e.(evMissionRedirected)
}

