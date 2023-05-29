package events

import (
	"time"
)

// Status event structure
type StatusT struct {
	Flags     uint64 `mapstructure:"Flags"`
	Flags2    uint64 `mapstructure:"Flags2"`
	Pips      [3]int `mapstructure:"Pips"`
	Firegroup int    `mapstructure:"Firegroup"`
	GuiFocus  int    `mapstructure:"GuiFocus"`
	Fuel      struct {
		FuelMain      float64 `mapstructure:"FuelMain"`
		FuelReservoir float64 `mapstructure:"FuelReservoir"`
	} `mapstructure:"Fuel"`
	Cargo        int     `mapstructure:"Cargo"`
	LegalState   string  `mapstructure:"LegalState"`
	Latitude     float64 `mapstructure:"Latitude"`
	Longitude    float64 `mapstructure:"Longitude"`
	Altitude     int     `mapstructure:"Altitude"`
	Heading      int     `mapstructure:"Heading"`
	BodyName     string  `mapstructure:"BodyName"`
	PlanetRadius float64 `mapstructure:"PlanetRadius"`
	Balance      int64   `mapstructure:"Balance"`
	Destination  struct {
		System string `mapstructure:"System"`
		Body   string `mapstructure:"Body"`
		Name   string `mapstructure:"Name"`
	} `mapstructure:"Destination"`
	Oxygen         float64   `mapstructure:"Oxygen"`
	Health         float64   `mapstructure:"Health"`
	Temperature    float64   `mapstructure:"Temperature"`
	SelectedWeapon string    `mapstructure:"SelectedWeapon"`
	Gravity        string    `mapstructure:"Gravity"`
	Event          string    `mapstructure:"event"`
	Timestamp      time.Time `mapstructure:"timestamp"`
}

// Status event handler
func (evh *EventHandler) Status(eventData map[string]interface{}) {
	/*
	   ev := new(StatusT)
	   mapstructure.Decode(eventData, ev)

	   	if ev.Gravity > 1.0 {
	   		alert("high gravity landing")
	   	}

	   	if flags | flagLowFuel {
	   		alert("fuel level low")
	   	}

	   	if flags | flagsOverHeating {
	   		alert("heat level high")
	   	}
	*/
}

// flags
const (
	flagLowFuel       uint64 = 524288
	flagusOverHeating uint64 = 1048576
	flagFSDJump       uint64 = 1073741824
)
