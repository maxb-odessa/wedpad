package events

import (
	"time"
)

// Friends event structure
type FriendsT struct {
	Name      string    `mapstructure:"Name"`
	Status    string    `mapstructure:"Status"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Friends event handler
func (evHandler EventHandler) Friends(eventData map[string]interface{}) {
    // ev := new(FriendsT)
    // mapstructure.Decode(eventData, ev)
}

