package events

import (
	"time"
)

// SAASignalsFound event structure
type SAASignalsFoundT struct {
	BodyID   int    `mapstructure:"BodyID"`
	BodyName string `mapstructure:"BodyName"`
	Genuses  []struct {
		Genus          string `mapstructure:"Genus"`
		GenusLocalised string `mapstructure:"Genus_Localised"`
	} `mapstructure:"Genuses"`
	Signals []struct {
		Count         int    `mapstructure:"Count"`
		Type          string `mapstructure:"Type"`
		TypeLocalised string `mapstructure:"Type_Localised"`
	} `mapstructure:"Signals"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// SAASignalsFound event handler
func (evh *EventHandler) SAASignalsFound(eventData map[string]interface{}) {
	/*
	   ev := new(SAASignalsFoundT)
	   mapstructure.Decode(eventData, ev)

	   cs := evh.CurrentSystem()
	   cs.AddPlanetSignals(ev)
	*/
}
