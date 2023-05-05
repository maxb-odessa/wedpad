package events

import (
	"time"
)

// MaterialTrade event structure
type evMaterialTrade struct {
	MarketID int `mapstructure:"MarketID"`
	Paid     struct {
		Category          string `mapstructure:"Category"`
		Material          string `mapstructure:"Material"`
		MaterialLocalised string `mapstructure:"Material_Localised"`
		Quantity          int    `mapstructure:"Quantity"`
	} `mapstructure:"Paid"`
	Received struct {
		Category          string `mapstructure:"Category"`
		Material          string `mapstructure:"Material"`
		MaterialLocalised string `mapstructure:"Material_Localised"`
		Quantity          int    `mapstructure:"Quantity"`
	} `mapstructure:"Received"`
	TraderType string    `mapstructure:"TraderType"`
	Event      string    `mapstructure:"event"`
	Timestamp  time.Time `mapstructure:"timestamp"`
}

// MaterialTrade event handler
func MaterialTrade(e interface{}) {
    // ev := e.(evMaterialTrade)
}

