package events

import (
	"time"
)

// MaterialCollected event structure
type evMaterialCollected struct {
	Category      string    `mapstructure:"Category"`
	Count         int       `mapstructure:"Count"`
	Name          string    `mapstructure:"Name"`
	NameLocalised string    `mapstructure:"Name_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// MaterialCollected event handler
func MaterialCollected(e interface{}) {
    // ev := e.(evMaterialCollected)
}

