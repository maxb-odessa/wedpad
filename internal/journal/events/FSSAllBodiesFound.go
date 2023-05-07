package events

import (
	"time"
)

// FSSAllBodiesFound event structure
type evFSSAllBodiesFound struct {
	Count         int       `mapstructure:"Count"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	SystemName    string    `mapstructure:"SystemName"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// FSSAllBodiesFound event handler
func (evHandler EventHandler) FSSAllBodiesFound(eventData map[string]interface{}) {
    // ev := new(evFSSAllBodiesFound)
    // mapstructure.Decode(eventData, ev)
}
