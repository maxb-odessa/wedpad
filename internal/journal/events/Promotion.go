package events

import (
	"time"
)

// Promotion event structure
type evPromotion struct {
	Exobiologist int       `mapstructure:"Exobiologist"`
	Explore      int       `mapstructure:"Explore"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// Promotion event handler
func (evHandler EventHandler) Promotion(eventData map[string]interface{}) {
    // ev := new(evPromotion)
    // mapstructure.Decode(eventData, ev)
}

