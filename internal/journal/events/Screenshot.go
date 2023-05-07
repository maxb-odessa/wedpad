package events

import (
	"time"
)

// Screenshot event structure
type ScreenshotT struct {
	Body      string    `mapstructure:"Body"`
	Filename  string    `mapstructure:"Filename"`
	Height    int       `mapstructure:"Height"`
	System    string    `mapstructure:"System"`
	Width     int       `mapstructure:"Width"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Screenshot event handler
func (evHandler EventHandler) Screenshot(eventData map[string]interface{}) {
    // ev := new(ScreenshotT)
    // mapstructure.Decode(eventData, ev)
}

