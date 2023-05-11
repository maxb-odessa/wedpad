package events

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

// ScanBaryCentre event structure
type ScanBaryCentreT struct {
	AscendingNode      float64   `mapstructure:"AscendingNode"`
	BodyID             int       `mapstructure:"BodyID"`
	Eccentricity       float64   `mapstructure:"Eccentricity"`
	MeanAnomaly        float64   `mapstructure:"MeanAnomaly"`
	OrbitalInclination float64   `mapstructure:"OrbitalInclination"`
	OrbitalPeriod      float64   `mapstructure:"OrbitalPeriod"`
	Periapsis          float64   `mapstructure:"Periapsis"`
	SemiMajorAxis      float64   `mapstructure:"SemiMajorAxis"`
	StarSystem         string    `mapstructure:"StarSystem"`
	SystemAddress      int       `mapstructure:"SystemAddress"`
	Event              string    `mapstructure:"event"`
	Timestamp          time.Time `mapstructure:"timestamp"`
}

// ScanBaryCentre event handler
func (evh *EventHandler) ScanBaryCentre(eventData map[string]interface{}) {
	ev := new(ScanBaryCentreT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()
	cs.SetName(ev.StarSystem)

	cs.AddBaryCentre(ev)
}
