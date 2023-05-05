package events

import (
	"time"
)

// AfmuRepairs event structure
type evAfmuRepairs struct {
	FullyRepaired   bool      `mapstructure:"FullyRepaired"`
	Health          float64   `mapstructure:"Health"`
	Module          string    `mapstructure:"Module"`
	ModuleLocalised string    `mapstructure:"Module_Localised"`
	Event           string    `mapstructure:"event"`
	Timestamp       time.Time `mapstructure:"timestamp"`
}

// AfmuRepairs event handler
func AfmuRepairs(e interface{}) {
    // ev := e.(evAfmuRepairs)
}

