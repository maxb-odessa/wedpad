package events

import (
	"time"
)

// FCMaterials event structure
type FCMaterialsT struct {
	CarrierID   string        `mapstructure:"CarrierID"`
	CarrierName string        `mapstructure:"CarrierName"`
	Items       []interface{} `mapstructure:"Items"`
	MarketID    int           `mapstructure:"MarketID"`
	Event       string        `mapstructure:"event"`
	Timestamp   time.Time     `mapstructure:"timestamp"`
}

// FCMaterials event handler
func (evh *EventHandler) FCMaterials(eventData map[string]interface{}) {
    // ev := new(FCMaterialsT)
    // mapstructure.Decode(eventData, ev)
}

