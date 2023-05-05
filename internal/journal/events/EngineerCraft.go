package events

import (
	"time"
)

// EngineerCraft event structure
type evEngineerCraft struct {
	ApplyExperimentalEffect     string `mapstructure:"ApplyExperimentalEffect"`
	BlueprintID                 int    `mapstructure:"BlueprintID"`
	BlueprintName               string `mapstructure:"BlueprintName"`
	Engineer                    string `mapstructure:"Engineer"`
	EngineerID                  int    `mapstructure:"EngineerID"`
	ExperimentalEffect          string `mapstructure:"ExperimentalEffect"`
	ExperimentalEffectLocalised string `mapstructure:"ExperimentalEffect_Localised"`
	Ingredients                 []struct {
		Count         int    `mapstructure:"Count"`
		Name          string `mapstructure:"Name"`
		NameLocalised string `mapstructure:"Name_Localised"`
	} `mapstructure:"Ingredients"`
	Level     int `mapstructure:"Level"`
	Modifiers []struct {
		Label         string  `mapstructure:"Label"`
		LessIsGood    int     `mapstructure:"LessIsGood"`
		OriginalValue float64 `mapstructure:"OriginalValue"`
		Value         float64 `mapstructure:"Value"`
	} `mapstructure:"Modifiers"`
	Module    string    `mapstructure:"Module"`
	Quality   float64   `mapstructure:"Quality"`
	Slot      string    `mapstructure:"Slot"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// EngineerCraft event handler
func EngineerCraft(e interface{}) {
    // ev := e.(evEngineerCraft)
}

