package events

import (
	"time"
)

// Backpack event structure
type BackpackT struct {
	Components  []interface{} `mapstructure:"Components"`
	Consumables []interface{} `mapstructure:"Consumables"`
	Data        []interface{} `mapstructure:"Data"`
	Items       []interface{} `mapstructure:"Items"`
	Event       string        `mapstructure:"event"`
	Timestamp   time.Time     `mapstructure:"timestamp"`
}

// Backpack event handler
func (evh *EventHandler) Backpack(eventData map[string]interface{}) {
    // ev := new(BackpackT)
    // mapstructure.Decode(eventData, ev)
}

