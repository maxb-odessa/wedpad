package events

import (
	"time"

	"github.com/maxb-odessa/slog"
	"github.com/mitchellh/mapstructure"
)

// Scan event structure
type ScanT struct {
	AbsoluteMagnitude     float64 `mapstructure:"AbsoluteMagnitude"`
	AgeMy                 int     `mapstructure:"Age_MY"`
	AscendingNode         float64 `mapstructure:"AscendingNode"`
	Atmosphere            string  `mapstructure:"Atmosphere"`
	AtmosphereComposition []struct {
		Name    string  `mapstructure:"Name"`
		Percent float64 `mapstructure:"Percent"`
	} `mapstructure:"AtmosphereComposition"`
	AtmosphereType string  `mapstructure:"AtmosphereType"`
	AxialTilt      float64 `mapstructure:"AxialTilt"`
	BodyID         int     `mapstructure:"BodyID"`
	BodyName       string  `mapstructure:"BodyName"`
	Composition    *struct {
		Ice   float64 `mapstructure:"Ice"`
		Metal float64 `mapstructure:"Metal"`
		Rock  float64 `mapstructure:"Rock"`
	} `mapstructure:"Composition"`
	DistanceFromArrivalLs float64 `mapstructure:"DistanceFromArrivalLS"`
	Eccentricity          float64 `mapstructure:"Eccentricity"`
	Landable              bool    `mapstructure:"Landable"`
	Luminosity            string  `mapstructure:"Luminosity"`
	MassEm                float64 `mapstructure:"MassEM"`
	Materials             []struct {
		Name    string  `mapstructure:"Name"`
		Percent float64 `mapstructure:"Percent"`
	} `mapstructure:"Materials"`
	MeanAnomaly        float64 `mapstructure:"MeanAnomaly"`
	OrbitalInclination float64 `mapstructure:"OrbitalInclination"`
	OrbitalPeriod      float64 `mapstructure:"OrbitalPeriod"`
	Parents            []struct {
		Null   int `mapstructure:"Null,omitempty"`
		Planet int `mapstructure:"Planet,omitempty"`
		Ring   int `mapstructure:"Ring,omitempty"`
		Star   int `mapstructure:"Star,omitempty"`
	} `mapstructure:"Parents"`
	Periapsis    float64 `mapstructure:"Periapsis"`
	PlanetClass  string  `mapstructure:"PlanetClass"`
	Radius       float64 `mapstructure:"Radius"`
	ReserveLevel string  `mapstructure:"ReserveLevel"`
	Rings        []struct {
		InnerRad  float64 `mapstructure:"InnerRad"`
		MassMt    float64 `mapstructure:"MassMT"`
		Name      string  `mapstructure:"Name"`
		OuterRad  float64 `mapstructure:"OuterRad"`
		RingClass string  `mapstructure:"RingClass"`
	} `mapstructure:"Rings"`
	RotationPeriod     float64   `mapstructure:"RotationPeriod"`
	ScanType           string    `mapstructure:"ScanType"`
	SemiMajorAxis      float64   `mapstructure:"SemiMajorAxis"`
	StarSystem         string    `mapstructure:"StarSystem"`
	StarType           string    `mapstructure:"StarType"`
	StellarMass        float64   `mapstructure:"StellarMass"`
	Subclass           int       `mapstructure:"Subclass"`
	SurfaceGravity     float64   `mapstructure:"SurfaceGravity"`
	SurfacePressure    float64   `mapstructure:"SurfacePressure"`
	SurfaceTemperature float64   `mapstructure:"SurfaceTemperature"`
	SystemAddress      int       `mapstructure:"SystemAddress"`
	TerraformState     string    `mapstructure:"TerraformState"`
	TidalLock          bool      `mapstructure:"TidalLock"`
	Volcanism          string    `mapstructure:"Volcanism"`
	WasDiscovered      bool      `mapstructure:"WasDiscovered"`
	WasMapped          bool      `mapstructure:"WasMapped"`
	Event              string    `mapstructure:"event"`
	Timestamp          time.Time `mapstructure:"timestamp"`
}

// Scan event handler
func (evh *EventHandler) Scan(eventData map[string]interface{}) {
	ev := new(ScanT)

	// init Parents with invalid values
	// this array is of a variable length, also not every field may be present
	// because Parent.Star (and other) could be zero (which is also a default value in Go)
	// mapstructure doesn't provide a way to set defaults yet
	ev.Parents = make([]struct {
		Null   int `mapstructure:"Null,omitempty"`
		Planet int `mapstructure:"Planet,omitempty"`
		Ring   int `mapstructure:"Ring,omitempty"`
		Star   int `mapstructure:"Star,omitempty"`
	}, 10) // TODO is 10 enuff?
	for i, _ := range ev.Parents {
		ev.Parents[i].Null = -1
		ev.Parents[i].Planet = -1
		ev.Parents[i].Ring = -1
		ev.Parents[i].Star = -1
	}

	mapstructure.Decode(eventData, ev)
	slog.Debug(9, "Parents for '%s': %+v", ev.BodyName, ev.Parents)
	cs := evh.CurrentSystem()

	cs.SetName(ev.StarSystem)

	if ev.StarType != "" {
		evh.scanStar(ev)
	} else if ev.PlanetClass != "" {
		evh.scanPlanet(ev)
	} else {
		slog.Debug(9, "Unknown 'Scan' type: not Star nor Planet (belt?), ignored")
	}
}

func (evh *EventHandler) scanStar(ev *ScanT) {
	cs := evh.CurrentSystem()
	cs.AddStar(ev)
	cs.ShowStars()
}

func (evh *EventHandler) scanPlanet(ev *ScanT) {
	cs := evh.CurrentSystem()
	cs.AddPlanet(ev)
	cs.ShowPlanets()
}
