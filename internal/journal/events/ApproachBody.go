package events

import (
	"time"
)

// ApproachBody event structure
type ApproachBodyT struct {
	Body          string    `mapstructure:"Body"`
	BodyID        int       `mapstructure:"BodyID"`
	StarSystem    string    `mapstructure:"StarSystem"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// ApproachBody event handler
func (evHandler EventHandler) ApproachBody(eventData map[string]interface{}) {
    // ev := new(ApproachBodyT)
    // mapstructure.Decode(eventData, ev)
}

