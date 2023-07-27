package events

import (
	"fmt"
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// CarrierJump event structure
type CarrierJumpT struct {
	Body     string `mapstructure:"Body"`
	BodyID   int    `mapstructure:"BodyID"`
	BodyType string `mapstructure:"BodyType"`
	Docked   bool   `mapstructure:"Docked"`
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
		RecoveringStates   []struct {
			State string `mapstructure:"State"`
			Trend int    `mapstructure:"Trend"`
		} `mapstructure:"RecoveringStates"`
	} `mapstructure:"Factions"`
	MarketID         int       `mapstructure:"MarketID"`
	Multicrew        bool      `mapstructure:"Multicrew"`
	OnFoot           bool      `mapstructure:"OnFoot"`
	Population       int       `mapstructure:"Population"`
	StarPos          []float64 `mapstructure:"StarPos"`
	StarSystem       string    `mapstructure:"StarSystem"`
	StationEconomies []struct {
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

// CarrierJump event handler
func (evh *EventHandler) CarrierJump(eventData map[string]interface{}) {
	ev := new(CarrierJumpT)
	mapstructure.Decode(eventData, ev)
	cs := evh.CurrentSystem()

	if !ev.Docked || ev.BodyType == "Star" {
		return
	}

	cs.SetName(ev.StarSystem)
	cs.SetMainStarID(ev.BodyID)
	cs.SetMainStarName(ev.Body)

	m := &msg.Message{
		Target: msg.TARGET_STARS,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_BODIES,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_SIGNALS,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_NOTES,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_STARS,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_BUTTON,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_BODIES,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_BUTTON,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_SIGNALS,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_BUTTON,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_NOTES,
		Action: msg.ACTION_CLEAN,
		Type:   msg.TYPE_BUTTON,
		Data:   "",
	}
	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_LOG,
		Action: msg.ACTION_APPEND,
		Type:   msg.TYPE_VIEW,
		Data: fmt.Sprintf("Carrier jumped in sytem <u><b>%s</b></u>", ev.StarSystem)
	}
	m.Send()

}
