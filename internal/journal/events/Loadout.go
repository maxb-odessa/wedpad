package events

import (
	"time"
)

// Loadout event structure
type evLoadout struct {
	CargoCapacity int `mapstructure:"CargoCapacity"`
	FuelCapacity  struct {
		Main    float64 `mapstructure:"Main"`
		Reserve float64 `mapstructure:"Reserve"`
	} `mapstructure:"FuelCapacity"`
	HullHealth   float64 `mapstructure:"HullHealth"`
	HullValue    int     `mapstructure:"HullValue"`
	MaxJumpRange float64 `mapstructure:"MaxJumpRange"`
	Modules      []struct {
		AmmoInClip   int `mapstructure:"AmmoInClip"`
		AmmoInHopper int `mapstructure:"AmmoInHopper"`
		Engineering  *struct {
			BlueprintID                 int    `mapstructure:"BlueprintID"`
			BlueprintName               string `mapstructure:"BlueprintName"`
			Engineer                    string `mapstructure:"Engineer"`
			EngineerID                  int    `mapstructure:"EngineerID"`
			ExperimentalEffect          string `mapstructure:"ExperimentalEffect"`
			ExperimentalEffectLocalised string `mapstructure:"ExperimentalEffect_Localised"`
			Level                       int    `mapstructure:"Level"`
			Modifiers                   []struct {
				Label         string  `mapstructure:"Label"`
				LessIsGood    int     `mapstructure:"LessIsGood"`
				OriginalValue float64 `mapstructure:"OriginalValue"`
				Value         float64 `mapstructure:"Value"`
			} `mapstructure:"Modifiers"`
			Quality float64 `mapstructure:"Quality"`
		} `mapstructure:"Engineering"`
		Health   float64 `mapstructure:"Health"`
		Item     string  `mapstructure:"Item"`
		On       bool    `mapstructure:"On"`
		Priority int     `mapstructure:"Priority"`
		Slot     string  `mapstructure:"Slot"`
		Value    int     `mapstructure:"Value"`
	} `mapstructure:"Modules"`
	ModulesValue int       `mapstructure:"ModulesValue"`
	Rebuy        int       `mapstructure:"Rebuy"`
	Ship         string    `mapstructure:"Ship"`
	ShipID       int       `mapstructure:"ShipID"`
	ShipIdent    string    `mapstructure:"ShipIdent"`
	ShipName     string    `mapstructure:"ShipName"`
	UnladenMass  float64   `mapstructure:"UnladenMass"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// Loadout event handler
func Loadout(e interface{}) {
    // ev := e.(evLoadout)
}

