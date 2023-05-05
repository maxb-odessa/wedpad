package events

import (
	"time"
)

// Friends event structure
type evFriends struct {
	Name      string    `mapstructure:"Name"`
	Status    string    `mapstructure:"Status"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Friends event handler
func Friends(eventData map[string]interface{}) {
    // ev := new(evFriends)
    // mapstructure.Decode(eventData, ev)
}

