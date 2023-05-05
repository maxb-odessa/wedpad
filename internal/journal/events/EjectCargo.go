package events

import (
	"time"
)

// EjectCargo event structure
type evEjectCargo struct {
	Abandoned     bool      `mapstructure:"Abandoned"`
	Count         int       `mapstructure:"Count"`
	Type          string    `mapstructure:"Type"`
	TypeLocalised string    `mapstructure:"Type_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// EjectCargo event handler
func EjectCargo(e interface{}) {
    // ev := e.(evEjectCargo)
}

