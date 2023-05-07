package events

import (
	"time"
)

// StoredModules event structure
type evStoredModules struct {
	Items []struct {
		BuyPrice              int     `mapstructure:"BuyPrice"`
		EngineerModifications string  `mapstructure:"EngineerModifications"`
		Hot                   bool    `mapstructure:"Hot"`
		InTransit             bool    `mapstructure:"InTransit"`
		Level                 int     `mapstructure:"Level"`
		MarketID              int     `mapstructure:"MarketID"`
		Name                  string  `mapstructure:"Name"`
		NameLocalised         string  `mapstructure:"Name_Localised"`
		Quality               float64 `mapstructure:"Quality"`
		StarSystem            string  `mapstructure:"StarSystem"`
		StorageSlot           int     `mapstructure:"StorageSlot"`
		TransferCost          int     `mapstructure:"TransferCost"`
		TransferTime          int     `mapstructure:"TransferTime"`
	} `mapstructure:"Items"`
	MarketID    int       `mapstructure:"MarketID"`
	StarSystem  string    `mapstructure:"StarSystem"`
	StationName string    `mapstructure:"StationName"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// StoredModules event handler
func (evHandler EventHandler) StoredModules(eventData map[string]interface{}) {
    // ev := new(evStoredModules)
    // mapstructure.Decode(eventData, ev)
}

