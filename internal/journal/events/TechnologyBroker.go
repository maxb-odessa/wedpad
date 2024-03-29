package events

import (
	"time"
)

// TechnologyBroker event structure
type TechnologyBrokerT struct {
	BrokerType    string        `mapstructure:"BrokerType"`
	Commodities   []interface{} `mapstructure:"Commodities"`
	ItemsUnlocked []struct {
		Name          string `mapstructure:"Name"`
		NameLocalised string `mapstructure:"Name_Localised"`
	} `mapstructure:"ItemsUnlocked"`
	MarketID  int `mapstructure:"MarketID"`
	Materials []struct {
		Category      string `mapstructure:"Category"`
		Count         int    `mapstructure:"Count"`
		Name          string `mapstructure:"Name"`
		NameLocalised string `mapstructure:"Name_Localised"`
	} `mapstructure:"Materials"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// TechnologyBroker event handler
func (evh *EventHandler) TechnologyBroker(eventData map[string]interface{}) {
    // ev := new(TechnologyBrokerT)
    // mapstructure.Decode(eventData, ev)
}

