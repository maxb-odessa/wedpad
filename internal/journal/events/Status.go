package events

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Status event structure
type StatusT struct {
	Flags     uint32 `mapstructure:"Flags"`
	Flags2    uint32 `mapstructure:"Flags2"`
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
		cs.alert.Alert("gravity", ALERT_LEVEL_WARN, fmt.Sprintf("High gravity: %.1fG!", ev.Gravity))
		cs.sound.Play("gravity")
	} else {
		cs.alert.Alert("gravity", ALERT_LEVEL_NONE, "")
	}

	if ev.Flags&flagOverHeating != 0 {
		cs.alert.Alert("heat", ALERT_LEVEL_WARN, "Overheating!")
		cs.sound.Play("heat")
	} else {
		cs.alert.Alert("heat", ALERT_LEVEL_NONE, "")
	}

	if ev.Flags&flagLowFuel != 0 {
		cs.sound.Play("fuel")
		cs.alert.Alert("fuel", ALERT_LEVEL_WARN, "Fuel Level is < 25%%!")
	} else {
		cs.alert.Alert("fuel", ALERT_LEVEL_NONE, "")
	}

}

// Flags bitmaps
const (
	flagDocked                    = 0x00000001 // (on a landing pad)
	flagLanded                    = 0x00000002 //(on planet surface)
	flagLandingGearDown           = 0x00000004
	flagShieldsUp                 = 0x00000008
	flagSupercruise               = 0x00000010
	flagFlightAssistOff           = 0x00000020
	flagHardpointsDeployed        = 0x00000040
	flagInWing                    = 0x00000080
	flagLightsOn                  = 0x00000100
	flagCargoScoopDeployed        = 0x00000200
	flagSilentRunning             = 0x00000400
	flagScoopingFuel              = 0x00000800
	flagSRVHandbrake              = 0x00001000
	flagSRVusingTurretView        = 0x00002000
	flagSRVTurretRetracted        = 0x00004000 //(close to ship)
	flagSRVDriveAssist            = 0x00008000
	flagFSDMassLocked             = 0x00010000
	flagFSDCharging               = 0x00020000
	flagFSDCooldown               = 0x00040000
	flagLowFuel                   = 0x00080000 // ( < 25% )
	flagOverHeating               = 0x00100000 // ( > 100% )
	flagHasLatLong                = 0x00200000
	flagIsInDanger                = 0x00400000
	flagBeingInterdicted          = 0x00800000
	flagInMainShip                = 0x01000000
	flagInFighter                 = 0x02000000
	flagInSRV                     = 0x04000000
	flagHudInAnalysisMode         = 0x08000000
	flagNightVision               = 0x10000000
	flagAltitudefromAverageRadius = 0x20000000
	flagFSDJump                   = 0x40000000
	flagSRVHighBeam               = 0x80000000
)

// Flags2 bitmaps
const (
	flag2OnFoot               = 0x00000001
	flag2InTaxi               = 0x00000002
	flag2InMulticrew          = 0x00000004
	flag2OnFootInStation      = 0x00000008
	flag2OnFootOnPlanet       = 0x00000010
	flag2AimDownSight         = 0x00000020
	flag2LowOxygen            = 0x00000040
	flag2LowHealth            = 0x00000080
	flag2Cold                 = 0x00000100
	flag2Hot                  = 0x00000200
	flag2VeryCold             = 0x00000400
	flag2VeryHot              = 0x00000800
	flag2Glide                = 0x00001000
	flag2OnFootInHangar       = 0x00002000
	flag2OnFootSocialSpace    = 0x00004000
	flag2OnFootExterior       = 0x00008000
	flag2BreathableAtmosphere = 0x00010000
	flag2Telepresence         = 0x00020000
	flag2Physical             = 0x00040000
	flag2FSD                  = 0x00080000
)
