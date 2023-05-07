package events

import (
	"time"
)

// FSSBodySignals event structure
type evFSSBodySignals struct {
	BodyID   int    `mapstructure:"BodyID"`
	BodyName string `mapstructure:"BodyName"`
	Signals  []struct {
		Count         int    `mapstructure:"Count"`
		Type          string `mapstructure:"Type"`
		TypeLocalised string `mapstructure:"Type_Localised"`
	} `mapstructure:"Signals"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// FSSBodySignals event handler
func (evHandler EventHandler) FSSBodySignals(eventData map[string]interface{}) {
    // ev := new(evFSSBodySignals)
    // mapstructure.Decode(eventData, ev)
}

