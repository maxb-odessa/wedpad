package events

import (
	"time"
)

// ModuleSellRemote event structure
type evModuleSellRemote struct {
	SellItem          string    `mapstructure:"SellItem"`
	SellItemLocalised string    `mapstructure:"SellItem_Localised"`
	SellPrice         int       `mapstructure:"SellPrice"`
	ServerID          int       `mapstructure:"ServerId"`
	Ship              string    `mapstructure:"Ship"`
	ShipID            int       `mapstructure:"ShipID"`
	StorageSlot       int       `mapstructure:"StorageSlot"`
	Event             string    `mapstructure:"event"`
	Timestamp         time.Time `mapstructure:"timestamp"`
}

// ModuleSellRemote event handler
func ModuleSellRemote(e interface{}) {
    // ev := e.(evModuleSellRemote)
}

