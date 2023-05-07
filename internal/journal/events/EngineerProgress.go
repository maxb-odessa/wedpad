package events

import (
	"time"
)

// EngineerProgress event structure
type EngineerProgressT struct {
	Engineer   string `mapstructure:"Engineer"`
	EngineerID int    `mapstructure:"EngineerID"`
	Engineers  []struct {
		Engineer     string `mapstructure:"Engineer"`
		EngineerID   int    `mapstructure:"EngineerID"`
		Progress     string `mapstructure:"Progress"`
		Rank         int    `mapstructure:"Rank"`
		RankProgress int    `mapstructure:"RankProgress"`
	} `mapstructure:"Engineers"`
	Rank      int       `mapstructure:"Rank"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// EngineerProgress event handler
func (evHandler EventHandler) EngineerProgress(eventData map[string]interface{}) {
    // ev := new(EngineerProgressT)
    // mapstructure.Decode(eventData, ev)
}

