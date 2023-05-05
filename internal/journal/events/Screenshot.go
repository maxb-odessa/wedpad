package events

import (
	"time"
)

// Screenshot event structure
type evScreenshot struct {
	Body      string    `mapstructure:"Body"`
	Filename  string    `mapstructure:"Filename"`
	Height    int       `mapstructure:"Height"`
	System    string    `mapstructure:"System"`
	Width     int       `mapstructure:"Width"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Screenshot event handler
func Screenshot(eventData map[string]interface{}) {
    // ev := new(evScreenshot)
    // mapstructure.Decode(eventData, ev)
}

