package events

import (
	"time"
)

// DockingRequested event structure
type DockingRequestedT struct {
	LandingPads struct {
		Large  int `mapstructure:"Large"`
		Medium int `mapstructure:"Medium"`
		Small  int `mapstructure:"Small"`
	} `mapstructure:"LandingPads"`
	MarketID    int       `mapstructure:"MarketID"`
	StationName string    `mapstructure:"StationName"`
	StationType string    `mapstructure:"StationType"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// DockingRequested event handler
func (evh *EventHandler) DockingRequested(eventData map[string]interface{}) {
    // ev := new(DockingRequestedT)
    // mapstructure.Decode(eventData, ev)
}

