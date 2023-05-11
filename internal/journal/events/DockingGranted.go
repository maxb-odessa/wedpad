package events

import (
	"time"
)

// DockingGranted event structure
type DockingGrantedT struct {
	LandingPad  int       `mapstructure:"LandingPad"`
	MarketID    int       `mapstructure:"MarketID"`
	StationName string    `mapstructure:"StationName"`
	StationType string    `mapstructure:"StationType"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// DockingGranted event handler
func (evh *EventHandler) DockingGranted(eventData map[string]interface{}) {
    // ev := new(DockingGrantedT)
    // mapstructure.Decode(eventData, ev)
}

