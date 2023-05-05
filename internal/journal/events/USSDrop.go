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
func USSDrop(e interface{}) {
    // ev := e.(evUSSDrop)
}

