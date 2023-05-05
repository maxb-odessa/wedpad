package events

import (
	"time"
)

// Fileheader event structure
type evFileheader struct {
	Odyssey     bool      `mapstructure:"Odyssey"`
	Build       string    `mapstructure:"build"`
	Event       string    `mapstructure:"event"`
	Gameversion string    `mapstructure:"gameversion"`
	Language    string    `mapstructure:"language"`
	Part        int       `mapstructure:"part"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// Fileheader event handler
func Fileheader(e interface{}) {
    // ev := e.(evFileheader)
}

