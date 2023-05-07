package events

import (
	"time"
)

// USSDrop event structure
type USSDropT struct {
	UssThreat        int       `mapstructure:"USSThreat"`
	UssType          string    `mapstructure:"USSType"`
	USSTypeLocalised string    `mapstructure:"USSType_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// USSDrop event handler
func (evHandler EventHandler) USSDrop(eventData map[string]interface{}) {
    // ev := new(USSDropT)
    // mapstructure.Decode(eventData, ev)
}

