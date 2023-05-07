package events

import (
	"time"
)

// BookTaxi event structure
type evBookTaxi struct {
	Cost                int       `mapstructure:"Cost"`
	DestinationLocation string    `mapstructure:"DestinationLocation"`
	DestinationSystem   string    `mapstructure:"DestinationSystem"`
	Event               string    `mapstructure:"event"`
	Timestamp           time.Time `mapstructure:"timestamp"`
}

// BookTaxi event handler
func (evHandler EventHandler) BookTaxi(eventData map[string]interface{}) {
    // ev := new(evBookTaxi)
    // mapstructure.Decode(eventData, ev)
}

