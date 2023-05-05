package events

import (
	"time"
)

// MaterialDiscovered event structure
type evMaterialDiscovered struct {
	Category        string    `mapstructure:"Category"`
	DiscoveryNumber int       `mapstructure:"DiscoveryNumber"`
	Name            string    `mapstructure:"Name"`
	NameLocalised   string    `mapstructure:"Name_Localised"`
	Event           string    `mapstructure:"event"`
	Timestamp       time.Time `mapstructure:"timestamp"`
}

// MaterialDiscovered event handler
func MaterialDiscovered(e interface{}) {
    // ev := e.(evMaterialDiscovered)
}

