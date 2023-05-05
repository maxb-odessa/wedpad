package events

import (
	"time"
)

// ModuleRetrieve event structure
type evModuleRetrieve struct {
	EngineerModifications  string    `mapstructure:"EngineerModifications"`
	Hot                    bool      `mapstructure:"Hot"`
	Level                  int       `mapstructure:"Level"`
	MarketID               int       `mapstructure:"MarketID"`
	Quality                float64   `mapstructure:"Quality"`
	RetrievedItem          string    `mapstructure:"RetrievedItem"`
	RetrievedItemLocalised string    `mapstructure:"RetrievedItem_Localised"`
	Ship                   string    `mapstructure:"Ship"`
	ShipID                 int       `mapstructure:"ShipID"`
	Slot                   string    `mapstructure:"Slot"`
	SwapOutItem            string    `mapstructure:"SwapOutItem"`
	SwapOutItemLocalised   string    `mapstructure:"SwapOutItem_Localised"`
	Event                  string    `mapstructure:"event"`
	Timestamp              time.Time `mapstructure:"timestamp"`
}

// ModuleRetrieve event handler
func ModuleRetrieve(eventData map[string]interface{}) {
    // ev := new(evModuleRetrieve)
    // mapstructure.Decode(eventData, ev)
}

