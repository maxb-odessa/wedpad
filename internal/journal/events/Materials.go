package events

import (
	"time"
)

// Materials event structure
type MaterialsT struct {
	Encoded []struct {
		Count         int    `mapstructure:"Count"`
		Name          string `mapstructure:"Name"`
		NameLocalised string `mapstructure:"Name_Localised"`
	} `mapstructure:"Encoded"`
	Manufactured []struct {
		Count         int    `mapstructure:"Count"`
		Name          string `mapstructure:"Name"`
		NameLocalised string `mapstructure:"Name_Localised"`
	} `mapstructure:"Manufactured"`
	Raw []struct {
		Count int    `mapstructure:"Count"`
		Name  string `mapstructure:"Name"`
	} `mapstructure:"Raw"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Materials event handler
func (evh *EventHandler) Materials(eventData map[string]interface{}) {
    // ev := new(MaterialsT)
    // mapstructure.Decode(eventData, ev)
}

