package events

import (
	"time"
)

// BookTaxi event structure
type BookTaxiT struct {
	Cost                int       `mapstructure:"Cost"`
	DestinationLocation string    `mapstructure:"DestinationLocation"`
	DestinationSystem   string    `mapstructure:"DestinationSystem"`
	Event               string    `mapstructure:"event"`
	Timestamp           time.Time `mapstructure:"timestamp"`
}

// BookTaxi event handler
func (evh *EventHandler) BookTaxi(eventData map[string]interface{}) {
    // ev := new(BookTaxiT)
    // mapstructure.Decode(eventData, ev)
}

