package events

import (
	"time"
)

// FCMaterials event structure
type evFCMaterials struct {
	CarrierID   string        `mapstructure:"CarrierID"`
	CarrierName string        `mapstructure:"CarrierName"`
	Items       []interface{} `mapstructure:"Items"`
	MarketID    int           `mapstructure:"MarketID"`
	Event       string        `mapstructure:"event"`
	Timestamp   time.Time     `mapstructure:"timestamp"`
}

// FCMaterials event handler
func FCMaterials(e interface{}) {
    // ev := e.(evFCMaterials)
}

