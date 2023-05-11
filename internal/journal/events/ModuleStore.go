package events

import (
	"time"
)

// ModuleStore event structure
type ModuleStoreT struct {
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
func (evh *EventHandler) ModuleStore(eventData map[string]interface{}) {
    // ev := new(ModuleStoreT)
    // mapstructure.Decode(eventData, ev)
}

