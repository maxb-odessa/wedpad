package events

import (
	"time"
)

// MissionCompleted event structure
type evMissionCompleted struct {
	Commodity          string `mapstructure:"Commodity"`
	CommodityLocalised string `mapstructure:"Commodity_Localised"`
	Count              int    `mapstructure:"Count"`
	DestinationSystem  string `mapstructure:"DestinationSystem"`
	Faction            string `mapstructure:"Faction"`
	FactionEffects     []struct {
		Effects []struct {
			Effect          string `mapstructure:"Effect"`
			EffectLocalised string `mapstructure:"Effect_Localised"`
			Trend           string `mapstructure:"Trend"`
		} `mapstructure:"Effects"`
		Faction   string `mapstructure:"Faction"`
		Influence []struct {
			Influence     string `mapstructure:"Influence"`
			SystemAddress int    `mapstructure:"SystemAddress"`
			Trend         string `mapstructure:"Trend"`
		} `mapstructure:"Influence"`
		Reputation      string `mapstructure:"Reputation"`
		ReputationTrend string `mapstructure:"ReputationTrend"`
	} `mapstructure:"FactionEffects"`
	MaterialsReward []struct {
		Category          string `mapstructure:"Category"`
		CategoryLocalised string `mapstructure:"Category_Localised"`
		Count             int    `mapstructure:"Count"`
		Name              string `mapstructure:"Name"`
		NameLocalised     string `mapstructure:"Name_Localised"`
	} `mapstructure:"MaterialsReward"`
	MissionID int       `mapstructure:"MissionID"`
	Name      string    `mapstructure:"Name"`
	Reward    int       `mapstructure:"Reward"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// MissionCompleted event handler
func MissionCompleted(e interface{}) {
    // ev := e.(evMissionCompleted)
}

