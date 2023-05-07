package events

import (
	"time"
)

// SellOrganicData event structure
type SellOrganicDataT struct {
	BioData []struct {
		Bonus            int    `mapstructure:"Bonus"`
		Genus            string `mapstructure:"Genus"`
		GenusLocalised   string `mapstructure:"Genus_Localised"`
		Species          string `mapstructure:"Species"`
		SpeciesLocalised string `mapstructure:"Species_Localised"`
		Value            int    `mapstructure:"Value"`
	} `mapstructure:"BioData"`
	MarketID  int       `mapstructure:"MarketID"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// SellOrganicData event handler
func (evHandler EventHandler) SellOrganicData(eventData map[string]interface{}) {
    // ev := new(SellOrganicDataT)
    // mapstructure.Decode(eventData, ev)
}

