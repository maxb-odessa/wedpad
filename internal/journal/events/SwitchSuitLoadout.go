package events

import (
	"time"
)

// SwitchSuitLoadout event structure
type SwitchSuitLoadoutT struct {
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

// SwitchSuitLoadout event handler
func (evh *EventHandler) SwitchSuitLoadout(eventData map[string]interface{}) {
    // ev := new(SwitchSuitLoadoutT)
    // mapstructure.Decode(eventData, ev)
}

