package events

import (
	"time"
)

// ApproachSettlement event structure
type evApproachSettlement struct {
	BodyID        int       `mapstructure:"BodyID"`
	BodyName      string    `mapstructure:"BodyName"`
	Latitude      float64   `mapstructure:"Latitude"`
	Longitude     float64   `mapstructure:"Longitude"`
	MarketID      int       `mapstructure:"MarketID"`
	Name          string    `mapstructure:"Name"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// ApproachSettlement event handler
func ApproachSettlement(eventData map[string]interface{}) {
    // ev := new(evApproachSettlement)
    // mapstructure.Decode(eventData, ev)
}

