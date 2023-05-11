package events

import (
	"time"
)

// SAAScanComplete event structure
type SAAScanCompleteT struct {
	BodyID           int       `mapstructure:"BodyID"`
	BodyName         string    `mapstructure:"BodyName"`
	EfficiencyTarget int       `mapstructure:"EfficiencyTarget"`
	ProbesUsed       int       `mapstructure:"ProbesUsed"`
	SystemAddress    int       `mapstructure:"SystemAddress"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// SAAScanComplete event handler
func (evh *EventHandler) SAAScanComplete(eventData map[string]interface{}) {
    // ev := new(SAAScanCompleteT)
    // mapstructure.Decode(eventData, ev)
}

