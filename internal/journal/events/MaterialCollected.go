package events

import (
	"time"
)

// MaterialCollected event structure
type MaterialCollectedT struct {
	Category      string    `mapstructure:"Category"`
	Count         int       `mapstructure:"Count"`
	Name          string    `mapstructure:"Name"`
	NameLocalised string    `mapstructure:"Name_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// MaterialCollected event handler
func (evHandler EventHandler) MaterialCollected(eventData map[string]interface{}) {
    // ev := new(MaterialCollectedT)
    // mapstructure.Decode(eventData, ev)
}

