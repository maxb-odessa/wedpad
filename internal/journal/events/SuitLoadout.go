package events

import (
	"time"
)

// SuitLoadout event structure
type SuitLoadoutT struct {
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
func (evh *EventHandler) SuitLoadout(eventData map[string]interface{}) {
    // ev := new(SuitLoadoutT)
    // mapstructure.Decode(eventData, ev)
}

