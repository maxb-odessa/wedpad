package events

import (
	"time"
)

// ShipTargeted event structure
type ShipTargetedT struct {
	Faction            string    `mapstructure:"Faction"`
	HullHealth         float64   `mapstructure:"HullHealth"`
	LegalStatus        string    `mapstructure:"LegalStatus"`
	PilotName          string    `mapstructure:"PilotName"`
	PilotNameLocalised string    `mapstructure:"PilotName_Localised"`
	PilotRank          string    `mapstructure:"PilotRank"`
	ScanStage          int       `mapstructure:"ScanStage"`
	ShieldHealth       float64   `mapstructure:"ShieldHealth"`
	Ship               string    `mapstructure:"Ship"`
	ShipLocalised      string    `mapstructure:"Ship_Localised"`
	SquadronID         string    `mapstructure:"SquadronID"`
	TargetLocked       bool      `mapstructure:"TargetLocked"`
	Event              string    `mapstructure:"event"`
	Timestamp          time.Time `mapstructure:"timestamp"`
}

// ShipTargeted event handler
func (evh *EventHandler) ShipTargeted(eventData map[string]interface{}) {
    // ev := new(ShipTargetedT)
    // mapstructure.Decode(eventData, ev)
}

