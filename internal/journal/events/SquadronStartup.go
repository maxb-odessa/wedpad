package events

import (
	"time"
)

// SquadronStartup event structure
type evSquadronStartup struct {
	CurrentRank  int       `mapstructure:"CurrentRank"`
	SquadronName string    `mapstructure:"SquadronName"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// SquadronStartup event handler
func SquadronStartup(e interface{}) {
    // ev := e.(evSquadronStartup)
}

