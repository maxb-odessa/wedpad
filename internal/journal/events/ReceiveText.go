package events

import (
	"time"
)

// ReceiveText event structure
type ReceiveTextT struct {
	Channel          string    `mapstructure:"Channel"`
	From             string    `mapstructure:"From"`
	FromLocalised    string    `mapstructure:"From_Localised"`
	Message          string    `mapstructure:"Message"`
	MessageLocalised string    `mapstructure:"Message_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// ReceiveText event handler
func (evh *EventHandler) ReceiveText(eventData map[string]interface{}) {
    // ev := new(ReceiveTextT)
    // mapstructure.Decode(eventData, ev)
}

