package events

import (
	"time"
)

// ShipLocker event structure
type evShipLocker struct {
	Components  []interface{} `mapstructure:"Components"`
	Consumables []struct {
		Count         int    `mapstructure:"Count"`
		Name          string `mapstructure:"Name"`
		NameLocalised string `mapstructure:"Name_Localised"`
		OwnerID       int    `mapstructure:"OwnerID"`
	} `mapstructure:"Consumables"`
	Data      []interface{} `mapstructure:"Data"`
	Items     []interface{} `mapstructure:"Items"`
	Event     string        `mapstructure:"event"`
	Timestamp time.Time     `mapstructure:"timestamp"`
}

// ShipLocker event handler
func ShipLocker(e interface{}) {
    // ev := e.(evShipLocker)
}

