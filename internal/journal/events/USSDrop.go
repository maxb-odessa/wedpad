package events

import (
	"time"
)

// USSDrop event structure
type evUSSDrop struct {
	UssThreat        int       `mapstructure:"USSThreat"`
	UssType          string    `mapstructure:"USSType"`
	USSTypeLocalised string    `mapstructure:"USSType_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// USSDrop event handler
func USSDrop(eventData map[string]interface{}) {
    // ev := new(evUSSDrop)
    // mapstructure.Decode(eventData, ev)
}

