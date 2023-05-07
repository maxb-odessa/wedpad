package events

import (
	"time"
)

// StoredShips event structure
type evStoredShips struct {
	MarketID  int `mapstructure:"MarketID"`
	ShipsHere []struct {
		Hot               bool   `mapstructure:"Hot"`
		Name              string `mapstructure:"Name"`
		ShipID            int    `mapstructure:"ShipID"`
		ShipType          string `mapstructure:"ShipType"`
		ShipTypeLocalised string `mapstructure:"ShipType_Localised"`
		Value             int    `mapstructure:"Value"`
	} `mapstructure:"ShipsHere"`
	ShipsRemote []struct {
		Hot               bool   `mapstructure:"Hot"`
		InTransit         bool   `mapstructure:"InTransit"`
		Name              string `mapstructure:"Name"`
		ShipID            int    `mapstructure:"ShipID"`
		ShipMarketID      int    `mapstructure:"ShipMarketID"`
		ShipType          string `mapstructure:"ShipType"`
		ShipTypeLocalised string `mapstructure:"ShipType_Localised"`
		StarSystem        string `mapstructure:"StarSystem"`
		TransferPrice     int    `mapstructure:"TransferPrice"`
		TransferTime      int    `mapstructure:"TransferTime"`
		Value             int    `mapstructure:"Value"`
	} `mapstructure:"ShipsRemote"`
	StarSystem  string    `mapstructure:"StarSystem"`
	StationName string    `mapstructure:"StationName"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// StoredShips event handler
func (evHandler EventHandler) StoredShips(eventData map[string]interface{}) {
    // ev := new(evStoredShips)
    // mapstructure.Decode(eventData, ev)
}

