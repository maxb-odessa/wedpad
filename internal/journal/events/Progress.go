package events

import (
	"time"
)

// Progress event structure
type ProgressT struct {
	Cqc          int       `mapstructure:"CQC"`
	Combat       int       `mapstructure:"Combat"`
	Empire       int       `mapstructure:"Empire"`
	Exobiologist int       `mapstructure:"Exobiologist"`
	Explore      int       `mapstructure:"Explore"`
	Federation   int       `mapstructure:"Federation"`
	Soldier      int       `mapstructure:"Soldier"`
	Trade        int       `mapstructure:"Trade"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// Progress event handler
func (evHandler EventHandler) Progress(eventData map[string]interface{}) {
    // ev := new(ProgressT)
    // mapstructure.Decode(eventData, ev)
}

