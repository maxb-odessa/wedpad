package events

import (
	"time"
)

// ModuleStore event structure
type evModuleStore struct {
	EngineerModifications string    `mapstructure:"EngineerModifications"`
	Hot                   bool      `mapstructure:"Hot"`
	Level                 int       `mapstructure:"Level"`
	MarketID              int       `mapstructure:"MarketID"`
	Quality               float64   `mapstructure:"Quality"`
	Ship                  string    `mapstructure:"Ship"`
	ShipID                int       `mapstructure:"ShipID"`
	Slot                  string    `mapstructure:"Slot"`
	StoredItem            string    `mapstructure:"StoredItem"`
	StoredItemLocalised   string    `mapstructure:"StoredItem_Localised"`
	Event                 string    `mapstructure:"event"`
	Timestamp             time.Time `mapstructure:"timestamp"`
}

// ModuleStore event handler
func ModuleStore(eventData map[string]interface{}) {
    // ev := new(evModuleStore)
    // mapstructure.Decode(eventData, ev)
}

