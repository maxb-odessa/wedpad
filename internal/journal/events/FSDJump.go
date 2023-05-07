package events

import (
	"time"
)

// FSDJump event structure
type evFSDJump struct {
	Body      string `mapstructure:"Body"`
	BodyID    int    `mapstructure:"BodyID"`
	BodyType  string `mapstructure:"BodyType"`
	BoostUsed int    `mapstructure:"BoostUsed"`
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
	Factions []struct {
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
	FuelLevel              float64   `mapstructure:"FuelLevel"`
	FuelUsed               float64   `mapstructure:"FuelUsed"`
	JumpDist               float64   `mapstructure:"JumpDist"`
	Multicrew              bool      `mapstructure:"Multicrew"`
	Population             int       `mapstructure:"Population"`
	StarPos                []float64 `mapstructure:"StarPos"`
	StarSystem             string    `mapstructure:"StarSystem"`
	SystemAddress          int       `mapstructure:"SystemAddress"`
	SystemAllegiance       string    `mapstructure:"SystemAllegiance"`
	SystemEconomy          string    `mapstructure:"SystemEconomy"`
	SystemEconomyLocalised string    `mapstructure:"SystemEconomy_Localised"`
	SystemFaction          *struct {
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

// FSDJump event handler
func (evHandler EventHandler) FSDJump(eventData map[string]interface{}) {
    // ev := new(evFSDJump)
    // mapstructure.Decode(eventData, ev)
}

