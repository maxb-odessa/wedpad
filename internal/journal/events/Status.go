package events

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
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
	Gravity        float64   `mapstructure:"Gravity"`
	Event          string    `mapstructure:"event"`
	Timestamp      time.Time `mapstructure:"timestamp"`
}

// Status event handler
func (evh *EventHandler) Status(eventData map[string]interface{}) {
	ev := new(StatusT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()

	if ev.Gravity > 1.0 {
		cs.alert.Alert("gravity", ALERT_LEVEL_WARN, fmt.Sprintf("High gravity: %.2f G"))
		cs.sound.Play("gravity")
	} else {
		cs.alert.Alert("gravity", ALERT_LEVEL_NONE, "")
	}

	if ev.Flags&flagOverHeating != 0 {
		cs.alert.Alert("heat", ALERT_LEVEL_WARN, fmt.Sprintf("Overheating!"))
		cs.sound.Play("heat")
	} else {
		cs.alert.Alert("heat", ALERT_LEVEL_NONE, "")
	}

	if ev.Fuel.FuelMain < 10.0 {
		alvl := ALERT_LEVEL_NONE
		if ev.Fuel.FuelMain < 5.0 {
			alvl = ALERT_LEVEL_CRIT
			cs.sound.Play("fuel")
		} else {
			alvl = ALERT_LEVEL_WARN
		}
		cs.sound.Play("fuel")
		cs.alert.Alert("fuel", alvl, fmt.Sprintf("Fuel level: %.1f tones", ev.Fuel.FuelMain))
	} else {
		cs.alert.Alert("fuel", ALERT_LEVEL_NONE, "")
	}

}

// flags
const (
	flagLowFuel     uint64 = 524288
	flagOverHeating uint64 = 1048576
	flagFSDJump     uint64 = 1073741824
)
