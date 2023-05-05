package events

import (
	"time"
)

// BuySuit event structure
type evBuySuit struct {
	Name          string    `mapstructure:"Name"`
	NameLocalised string    `mapstructure:"Name_Localised"`
	Price         int       `mapstructure:"Price"`
	SuitID        int       `mapstructure:"SuitID"`
	SuitMods      []string  `mapstructure:"SuitMods"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// BuySuit event handler
func BuySuit(e interface{}) {
    // ev := e.(evBuySuit)
}

