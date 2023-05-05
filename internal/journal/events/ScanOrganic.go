package events

import (
	"time"
)

// ScanOrganic event structure
type evScanOrganic struct {
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
func ScanOrganic(e interface{}) {
    // ev := e.(evScanOrganic)
}

