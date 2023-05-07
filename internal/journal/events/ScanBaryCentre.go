package events

import (
	"time"
)

// ScanBaryCentre event structure
type evScanBaryCentre struct {
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
func (evHandler EventHandler) ScanBaryCentre(eventData map[string]interface{}) {
    // ev := new(evScanBaryCentre)
    // mapstructure.Decode(eventData, ev)
}

