package events

import (
	"time"
)

// CreateSuitLoadout event structure
type CreateSuitLoadoutT struct {
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

// CreateSuitLoadout event handler
func (evh *EventHandler) CreateSuitLoadout(eventData map[string]interface{}) {
    // ev := new(CreateSuitLoadoutT)
    // mapstructure.Decode(eventData, ev)
}

