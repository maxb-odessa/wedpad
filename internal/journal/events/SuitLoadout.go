package events

import (
	"time"
)

// SuitLoadout event structure
type evSuitLoadout struct {
	LoadoutID         int           `mapstructure:"LoadoutID"`
	LoadoutName       string        `mapstructure:"LoadoutName"`
	Modules           []interface{} `mapstructure:"Modules"`
	SuitID            int           `mapstructure:"SuitID"`
	SuitMods          []string      `mapstructure:"SuitMods"`
	SuitName          string        `mapstructure:"SuitName"`
	SuitNameLocalised string        `mapstructure:"SuitName_Localised"`
	Event             string        `mapstructure:"event"`
	Timestamp         time.Time     `mapstructure:"timestamp"`
}

// SuitLoadout event handler
func SuitLoadout(eventData map[string]interface{}) {
    // ev := new(evSuitLoadout)
    // mapstructure.Decode(eventData, ev)
}

