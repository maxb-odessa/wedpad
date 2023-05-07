package events

import (
	"time"
)

// SAAScanComplete event structure
type evSAAScanComplete struct {
	BodyID           int       `mapstructure:"BodyID"`
	BodyName         string    `mapstructure:"BodyName"`
	EfficiencyTarget int       `mapstructure:"EfficiencyTarget"`
	ProbesUsed       int       `mapstructure:"ProbesUsed"`
	SystemAddress    int       `mapstructure:"SystemAddress"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// SAAScanComplete event handler
func (evHandler EventHandler) SAAScanComplete(eventData map[string]interface{}) {
    // ev := new(evSAAScanComplete)
    // mapstructure.Decode(eventData, ev)
}

