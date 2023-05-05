package events

import (
	"time"
)

// Synthesis event structure
type evSynthesis struct {
	Materials []struct {
		Count         int    `mapstructure:"Count"`
		Name          string `mapstructure:"Name"`
		NameLocalised string `mapstructure:"Name_Localised"`
	} `mapstructure:"Materials"`
	Name      string    `mapstructure:"Name"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Synthesis event handler
func Synthesis(e interface{}) {
    // ev := e.(evSynthesis)
}

