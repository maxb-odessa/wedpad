package events

import (
	"time"
)

// ReceiveText event structure
type evReceiveText struct {
	Channel          string    `mapstructure:"Channel"`
	From             string    `mapstructure:"From"`
	FromLocalised    string    `mapstructure:"From_Localised"`
	Message          string    `mapstructure:"Message"`
	MessageLocalised string    `mapstructure:"Message_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// ReceiveText event handler
func ReceiveText(e interface{}) {
    // ev := e.(evReceiveText)
}

