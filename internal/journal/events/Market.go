package events

import (
	"time"
)

// Market event structure
type evMarket struct {
	CarrierDockingAccess string `mapstructure:"CarrierDockingAccess"`
	Items                []struct {
		BuyPrice          int    `mapstructure:"BuyPrice"`
		Category          string `mapstructure:"Category"`
		CategoryLocalised string `mapstructure:"Category_Localised"`
		Consumer          bool   `mapstructure:"Consumer"`
		Demand            int    `mapstructure:"Demand"`
		DemandBracket     int    `mapstructure:"DemandBracket"`
		MeanPrice         int    `mapstructure:"MeanPrice"`
		Name              string `mapstructure:"Name"`
		NameLocalised     string `mapstructure:"Name_Localised"`
		Producer          bool   `mapstructure:"Producer"`
		Rare              bool   `mapstructure:"Rare"`
		SellPrice         int    `mapstructure:"SellPrice"`
		Stock             int    `mapstructure:"Stock"`
		StockBracket      int    `mapstructure:"StockBracket"`
		ID                int    `mapstructure:"id"`
	} `mapstructure:"Items"`
	MarketID    int       `mapstructure:"MarketID"`
	StarSystem  string    `mapstructure:"StarSystem"`
	StationName string    `mapstructure:"StationName"`
	StationType string    `mapstructure:"StationType"`
	Event       string    `mapstructure:"event"`
	Timestamp   time.Time `mapstructure:"timestamp"`
}

// Market event handler
func Market(eventData map[string]interface{}) {
    // ev := new(evMarket)
    // mapstructure.Decode(eventData, ev)
}

