package events

import (
	"time"
)

// LoadGame event structure
type evLoadGame struct {
	Commander     string    `mapstructure:"Commander"`
	Credits       int       `mapstructure:"Credits"`
	Fid           string    `mapstructure:"FID"`
	FuelCapacity  float64   `mapstructure:"FuelCapacity"`
	FuelLevel     float64   `mapstructure:"FuelLevel"`
	GameMode      string    `mapstructure:"GameMode"`
	Horizons      bool      `mapstructure:"Horizons"`
	Loan          int       `mapstructure:"Loan"`
	Odyssey       bool      `mapstructure:"Odyssey"`
	Ship          string    `mapstructure:"Ship"`
	ShipID        int       `mapstructure:"ShipID"`
	ShipIdent     string    `mapstructure:"ShipIdent"`
	ShipName      string    `mapstructure:"ShipName"`
	ShipLocalised string    `mapstructure:"Ship_Localised"`
	StartLanded   bool      `mapstructure:"StartLanded"`
	Build         string    `mapstructure:"build"`
	Event         string    `mapstructure:"event"`
	Gameversion   string    `mapstructure:"gameversion"`
	Language      string    `mapstructure:"language"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// LoadGame event handler
func LoadGame(e interface{}) {
    // ev := e.(evLoadGame)
}

