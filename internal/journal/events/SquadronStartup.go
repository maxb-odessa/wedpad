package events

import (
	"time"
)

// SquadronStartup event structure
type SquadronStartupT struct {
	CurrentRank  int       `mapstructure:"CurrentRank"`
	SquadronName string    `mapstructure:"SquadronName"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// SquadronStartup event handler
func (evHandler EventHandler) SquadronStartup(eventData map[string]interface{}) {
    // ev := new(SquadronStartupT)
    // mapstructure.Decode(eventData, ev)
}

