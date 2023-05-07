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
func (evHandler EventHandler) Synthesis(eventData map[string]interface{}) {
    // ev := new(evSynthesis)
    // mapstructure.Decode(eventData, ev)
}
