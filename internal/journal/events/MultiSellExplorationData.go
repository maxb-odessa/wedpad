package events

import (
	"time"
)

// MultiSellExplorationData event structure
type MultiSellExplorationDataT struct {
	BaseValue  int `mapstructure:"BaseValue"`
	Bonus      int `mapstructure:"Bonus"`
	Discovered []struct {
		NumBodies           int    `mapstructure:"NumBodies"`
		SystemName          string `mapstructure:"SystemName"`
		SystemNameLocalised string `mapstructure:"SystemName_Localised"`
	} `mapstructure:"Discovered"`
	TotalEarnings int       `mapstructure:"TotalEarnings"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// MultiSellExplorationData event handler
func (evHandler EventHandler) MultiSellExplorationData(eventData map[string]interface{}) {
    // ev := new(MultiSellExplorationDataT)
    // mapstructure.Decode(eventData, ev)
}

