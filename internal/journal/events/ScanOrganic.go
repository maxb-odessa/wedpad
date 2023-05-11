package events

import (
	"time"
)

// ScanOrganic event structure
type ScanOrganicT struct {
	Body             int       `mapstructure:"Body"`
	Genus            string    `mapstructure:"Genus"`
	GenusLocalised   string    `mapstructure:"Genus_Localised"`
	ScanType         string    `mapstructure:"ScanType"`
	Species          string    `mapstructure:"Species"`
	SpeciesLocalised string    `mapstructure:"Species_Localised"`
	SystemAddress    int       `mapstructure:"SystemAddress"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// ScanOrganic event handler
func (evh *EventHandler) ScanOrganic(eventData map[string]interface{}) {
    // ev := new(ScanOrganicT)
    // mapstructure.Decode(eventData, ev)
}

