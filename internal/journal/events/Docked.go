package events

import (
	"time"
)

// Docked event structure
type DockedT struct {
	DistFromStarLs float64 `mapstructure:"DistFromStarLS"`
	LandingPads    struct {
		Large  int `mapstructure:"Large"`
		Medium int `mapstructure:"Medium"`
		Small  int `mapstructure:"Small"`
	} `mapstructure:"LandingPads"`
	MarketID          int    `mapstructure:"MarketID"`
	Multicrew         bool   `mapstructure:"Multicrew"`
	StarSystem        string `mapstructure:"StarSystem"`
	StationAllegiance string `mapstructure:"StationAllegiance"`
	StationEconomies  []struct {
		Name          string  `mapstructure:"Name"`
		NameLocalised string  `mapstructure:"Name_Localised"`
		Proportion    float64 `mapstructure:"Proportion"`
	} `mapstructure:"StationEconomies"`
	StationEconomy          string `mapstructure:"StationEconomy"`
	StationEconomyLocalised string `mapstructure:"StationEconomy_Localised"`
	StationFaction          struct {
		FactionState string `mapstructure:"FactionState"`
		Name         string `mapstructure:"Name"`
	} `mapstructure:"StationFaction"`
	StationGovernment          string    `mapstructure:"StationGovernment"`
	StationGovernmentLocalised string    `mapstructure:"StationGovernment_Localised"`
	StationName                string    `mapstructure:"StationName"`
	StationServices            []string  `mapstructure:"StationServices"`
	StationType                string    `mapstructure:"StationType"`
	SystemAddress              int       `mapstructure:"SystemAddress"`
	Taxi                       bool      `mapstructure:"Taxi"`
	Event                      string    `mapstructure:"event"`
	Timestamp                  time.Time `mapstructure:"timestamp"`
}

// Docked event handler
func (evHandler EventHandler) Docked(eventData map[string]interface{}) {
    // ev := new(DockedT)
    // mapstructure.Decode(eventData, ev)
}

