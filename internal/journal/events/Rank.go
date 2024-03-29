package events

import (
	"time"
)

// Rank event structure
type RankT struct {
	Cqc          int       `mapstructure:"CQC"`
	Combat       int       `mapstructure:"Combat"`
	Empire       int       `mapstructure:"Empire"`
	Exobiologist int       `mapstructure:"Exobiologist"`
	Explore      int       `mapstructure:"Explore"`
	Federation   int       `mapstructure:"Federation"`
	Soldier      int       `mapstructure:"Soldier"`
	Trade        int       `mapstructure:"Trade"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// Rank event handler
func (evh *EventHandler) Rank(eventData map[string]interface{}) {
    // ev := new(RankT)
    // mapstructure.Decode(eventData, ev)
}

