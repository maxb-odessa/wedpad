package events

import (
	"time"
)

// CargoTransfer event structure
type evCargoTransfer struct {
	Transfers []struct {
		Count     int    `mapstructure:"Count"`
		Direction string `mapstructure:"Direction"`
		Type      string `mapstructure:"Type"`
	} `mapstructure:"Transfers"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CargoTransfer event handler
func CargoTransfer(e interface{}) {
    // ev := e.(evCargoTransfer)
}

