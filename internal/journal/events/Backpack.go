package events

import (
	"time"
)

// Backpack event structure
type evBackpack struct {
	Components  []interface{} `mapstructure:"Components"`
	Consumables []interface{} `mapstructure:"Consumables"`
	Data        []interface{} `mapstructure:"Data"`
	Items       []interface{} `mapstructure:"Items"`
	Event       string        `mapstructure:"event"`
	Timestamp   time.Time     `mapstructure:"timestamp"`
}

// Backpack event handler
func Backpack(e interface{}) {
    // ev := e.(evBackpack)
}

