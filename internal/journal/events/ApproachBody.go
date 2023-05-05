package events

import (
	"time"
)

// ApproachBody event structure
type evApproachBody struct {
	Body          string    `mapstructure:"Body"`
	BodyID        int       `mapstructure:"BodyID"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// ApproachBody event handler
func ApproachBody(eventData map[string]interface{}) {
    // ev := new(evApproachBody)
    // mapstructure.Decode(eventData, ev)
}

