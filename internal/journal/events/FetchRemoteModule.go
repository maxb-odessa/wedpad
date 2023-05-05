package events

import (
	"time"
)

// FetchRemoteModule event structure
type evFetchRemoteModule struct {
	ServerID            int       `mapstructure:"ServerId"`
	Ship                string    `mapstructure:"Ship"`
	ShipID              int       `mapstructure:"ShipID"`
	StorageSlot         int       `mapstructure:"StorageSlot"`
	StoredItem          string    `mapstructure:"StoredItem"`
	StoredItemLocalised string    `mapstructure:"StoredItem_Localised"`
	TransferCost        int       `mapstructure:"TransferCost"`
	TransferTime        int       `mapstructure:"TransferTime"`
	Event               string    `mapstructure:"event"`
	Timestamp           time.Time `mapstructure:"timestamp"`
}

// FetchRemoteModule event handler
func FetchRemoteModule(e interface{}) {
    // ev := e.(evFetchRemoteModule)
}

