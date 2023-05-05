package events

import (
	"time"
)

// Location event structure
type evLocation struct {
	Body      string `mapstructure:"Body"`
	BodyID    int    `mapstructure:"BodyID"`
	BodyType  string `mapstructure:"BodyType"`
	Conflicts []struct {
		Faction1 struct {
			Name    string `mapstructure:"Name"`
			Stake   string `mapstructure:"Stake"`
			WonDays int    `mapstructure:"WonDays"`
		} `mapstructure:"Faction1"`
		Faction2 struct {
			Name    string `mapstructure:"Name"`
			Stake   string `mapstructure:"Stake"`
			WonDays int    `mapstructure:"WonDays"`
		} `mapstructure:"Faction2"`
		Status  string `mapstructure:"Status"`
		WarType string `mapstructure:"WarType"`
	} `mapstructure:"Conflicts"`
	DistFromStarLs float64 `mapstructure:"DistFromStarLS"`
	Docked         bool    `mapstructure:"Docked"`
	Factions       []struct {
		ActiveStates []struct {
			State string `mapstructure:"State"`
		} `mapstructure:"ActiveStates"`
		Allegiance         string  `mapstructure:"Allegiance"`
		FactionState       string  `mapstructure:"FactionState"`
		Government         string  `mapstructure:"Government"`
		Happiness          string  `mapstructure:"Happiness"`
		HappinessLocalised string  `mapstructure:"Happiness_Localised"`
		Influence          float64 `mapstructure:"Influence"`
		MyReputation       float64 `mapstructure:"MyReputation"`
		Name               string  `mapstructure:"Name"`
		PendingStates      []struct {
			State string `mapstructure:"State"`
			Trend int    `mapstructure:"Trend"`
		} `mapstructure:"PendingStates"`
		RecoveringStates []struct {
			State string `mapstructure:"State"`
			Trend int    `mapstructure:"Trend"`
		} `mapstructure:"RecoveringStates"`
	} `mapstructure:"Factions"`
	InSrv             bool      `mapstructure:"InSRV"`
	Latitude          float64   `mapstructure:"Latitude"`
	Longitude         float64   `mapstructure:"Longitude"`
	MarketID          int       `mapstructure:"MarketID"`
	Multicrew         bool      `mapstructure:"Multicrew"`
	OnFoot            bool      `mapstructure:"OnFoot"`
	Population        int       `mapstructure:"Population"`
	StarPos           []float64 `mapstructure:"StarPos"`
	StarSystem        string    `mapstructure:"StarSystem"`
	StationAllegiance string    `mapstructure:"StationAllegiance"`
	StationEconomies  []struct {
		Name          string  `mapstructure:"Name"`
		NameLocalised string  `mapstructure:"Name_Localised"`
		Proportion    float64 `mapstructure:"Proportion"`
	} `mapstructure:"StationEconomies"`
	StationEconomy          string `mapstructure:"StationEconomy"`
	StationEconomyLocalised string `mapstructure:"StationEconomy_Localised"`
	StationFaction          *struct {
		Name string `mapstructure:"Name"`
	} `mapstructure:"StationFaction"`
	StationGovernment          string   `mapstructure:"StationGovernment"`
	StationGovernmentLocalised string   `mapstructure:"StationGovernment_Localised"`
	StationName                string   `mapstructure:"StationName"`
	StationServices            []string `mapstructure:"StationServices"`
	StationType                string   `mapstructure:"StationType"`
	SystemAddress              int      `mapstructure:"SystemAddress"`
	SystemAllegiance           string   `mapstructure:"SystemAllegiance"`
	SystemEconomy              string   `mapstructure:"SystemEconomy"`
	SystemEconomyLocalised     string   `mapstructure:"SystemEconomy_Localised"`
	SystemFaction              *struct {
		FactionState string `mapstructure:"FactionState"`
		Name         string `mapstructure:"Name"`
	} `mapstructure:"SystemFaction"`
	SystemGovernment             string    `mapstructure:"SystemGovernment"`
	SystemGovernmentLocalised    string    `mapstructure:"SystemGovernment_Localised"`
	SystemSecondEconomy          string    `mapstructure:"SystemSecondEconomy"`
	SystemSecondEconomyLocalised string    `mapstructure:"SystemSecondEconomy_Localised"`
	SystemSecurity               string    `mapstructure:"SystemSecurity"`
	SystemSecurityLocalised      string    `mapstructure:"SystemSecurity_Localised"`
	Taxi                         bool      `mapstructure:"Taxi"`
	Event                        string    `mapstructure:"event"`
	Timestamp                    time.Time `mapstructure:"timestamp"`
}

// Location event handler
func Location(eventData map[string]interface{}) {
    // ev := new(evLocation)
    // mapstructure.Decode(eventData, ev)
}

