package events

import (
	"time"
)

// Interdicted event structure
type evInterdicted struct {
	Faction     string    `mapstructure:"Faction"`
	Interdictor string    `mapstructure:"Interdictor"`
	IsPlayer    bool      `mapstructure:"IsPlayer"`
	Submitted   bool      `mapstructure:"Submitted"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// Interdicted event handler
func Interdicted(e interface{}) {
    // ev := e.(evInterdicted)
}

