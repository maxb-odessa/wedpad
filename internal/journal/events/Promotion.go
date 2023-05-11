package events

import (
	"time"
)

// Promotion event structure
type PromotionT struct {
	Exobiologist int       `mapstructure:"Exobiologist"`
	Explore      int       `mapstructure:"Explore"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// Promotion event handler
func (evh *EventHandler) Promotion(eventData map[string]interface{}) {
    // ev := new(PromotionT)
    // mapstructure.Decode(eventData, ev)
}

