package events

import (
	"time"
)

// MissionAccepted event structure
type MissionAcceptedT struct {
	Commodity          string    `mapstructure:"Commodity"`
	CommodityLocalised string    `mapstructure:"Commodity_Localised"`
	Count              int       `mapstructure:"Count"`
	DestinationSystem  string    `mapstructure:"DestinationSystem"`
	Expiry             time.Time `mapstructure:"Expiry"`
	Faction            string    `mapstructure:"Faction"`
	Influence          string    `mapstructure:"Influence"`
	LocalisedName      string    `mapstructure:"LocalisedName"`
	MissionID          int       `mapstructure:"MissionID"`
	Name               string    `mapstructure:"Name"`
	PassengerCount     int       `mapstructure:"PassengerCount"`
	PassengerType      string    `mapstructure:"PassengerType"`
	PassengerViPs      bool      `mapstructure:"PassengerVIPs"`
	PassengerWanted    bool      `mapstructure:"PassengerWanted"`
	Reputation         string    `mapstructure:"Reputation"`
	Reward             int       `mapstructure:"Reward"`
	Wing               bool      `mapstructure:"Wing"`
	Event              string    `mapstructure:"event"`
	Timestamp          time.Time `mapstructure:"timestamp"`
}

// MissionAccepted event handler
func (evHandler EventHandler) MissionAccepted(eventData map[string]interface{}) {
    // ev := new(MissionAcceptedT)
    // mapstructure.Decode(eventData, ev)
}

