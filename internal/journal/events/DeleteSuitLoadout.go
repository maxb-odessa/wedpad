package events

import (
	"time"
)

// DeleteSuitLoadout event structure
type DeleteSuitLoadoutT struct {
	LoadoutID         int       `mapstructure:"LoadoutID"`
	LoadoutName       string    `mapstructure:"LoadoutName"`
	SuitID            int       `mapstructure:"SuitID"`
	SuitName          string    `mapstructure:"SuitName"`
	SuitNameLocalised string    `mapstructure:"SuitName_Localised"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// DeleteSuitLoadout event handler
func (evHandler EventHandler) DeleteSuitLoadout(eventData map[string]interface{}) {
    // ev := new(DeleteSuitLoadoutT)
    // mapstructure.Decode(eventData, ev)
}

