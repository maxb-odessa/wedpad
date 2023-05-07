package events

import (
	"fmt"
	"strings"
	"wedpad/internal/msg"
)

type CurrentSystemT struct {
	Name          string
	Sector        string
	Stars         map[int]*ScanT
	Planets       map[int]*ScanT
	BaryCentres   map[int]*ScanBaryCentreT
	Signals       map[int]*FSSSignalDiscoveredT
	PlanetSignals map[int]*SAASignalsFoundT
}

var CurrentSystem *CurrentSystemT

// TODO name, sector, etc
func (cs *CurrentSystemT) Clean(what string) {
	switch strings.ToLower(what) {
	case "stars":
		cs.Stars = make(map[int]*ScanT)
	case "planets":
		cs.Planets = make(map[int]*ScanT)
	case "barycentres":
		cs.BaryCentres = make(map[int]*ScanBaryCentreT)
	case "signals":
		cs.Signals = make(map[int]*FSSSignalDiscoveredT)
	case "planetSignals":
		cs.PlanetSignals = make(map[int]*SAASignalsFoundT)
	default:
		cs.Stars = make(map[int]*ScanT)
		cs.Planets = make(map[int]*ScanT)
		cs.BaryCentres = make(map[int]*ScanBaryCentreT)
		cs.Signals = make(map[int]*FSSSignalDiscoveredT)
		cs.PlanetSignals = make(map[int]*SAASignalsFoundT)
	}
}

func (cs *CurrentSystemT) ShowStars() {

	stars := make([]map[string]interface{}, 1)

	var keys []int
	for k, _ := range CurrentSystem.Stars {
		keys = append(keys, k)
	}

	for _, id := range keys {
		s := CurrentSystem.Stars[id]

		star := make(map[string]interface{})

		star["Id"] = s.BodyID
		star["Type"] = StarTypes[s.StarType]
		star["Subclass"] = s.Subclass
		star["Luminosity"] = s.Luminosity
		star["DistanceLs"] = Num(s.DistanceFromArrivalLs)
		star["RadiusS"] = Num(s.Radius / SOLAR_RADIUS)
		star["MassS"] = Num(s.StellarMass)
		star["TemperatureK"] = Num(s.SurfaceTemperature)
		star["Discovered"] = WasDiscovered[s.WasDiscovered]
		rn, rr := CalcRings(s)
		if rn > 0 {
			star["RingsNum"] = rn
			star["RingsRadiusLs"] = Num(rr)
		}

		stars = append(stars, star)
	}

	m := &msg.Message{
		Action: msg.ACTION_REPLACE,
		Target: msg.TARGET_SYSTEM,
		Type:   msg.TYPE_VIEW,
		Data:   stars,
	}

	m.Send()
}

// in meters
const (
	SOLAR_RADIUS    = 696340000.0
	EARTH_RADIUS    = 6371.0 * 1000.0
	LIGHT_SECOND    = 299792.0 * 1000.0
	SECONDS_IN_HOUR = 60.0 * 60.0
	SECONDS_IN_DAY  = 24.0 * 60.0 * 60.0
)

func Num(val float64) (s string) {

	switch {
	case val >= 1_000_000_000.0:
		s = fmt.Sprintf("%3.1fG", val/1_000_000_000.0)
	case val >= 1_000_000.0:
		s = fmt.Sprintf("%3.1fM", val/1_000_000.0)
	case val >= 1_000.0:
		s = fmt.Sprintf("%3.1fK", val/1_000.0)
	case val >= 1:
		s = fmt.Sprintf("%3.1f", val)
	default:
		s = fmt.Sprintf("%.2f", val)
	}

	return
}

func CalcRings(ev *ScanT) (num int, rad float64) {

	for _, r := range ev.Rings {
		if r.Name[len(r.Name)-5:] != " Ring" {
			continue
		}
		num++
		if rad < r.OuterRad {
			rad = r.OuterRad
		}
	}

	return
}

var WasDiscovered = map[bool]StrColPair{
	true:  {"Y", "#ffff00"},
	false: {"N", "#0000ff"},
}

type StrColPair struct {
	Str, Color string
}

var StarTypes = map[string]StrColPair{
	"A":                     {"A", "#ffaaff"},
	"A_BlueWhiteSuperGiant": {"A++", "#ffaaff"},
	"AeBe":                  {"AeBe", "#ffaaff"},
	"B":                     {"B", "#ffaaff"},
	"B_BlueWhiteSuperGiant": {"B++", "#ffaaff"},
	"C":                     {"C", "#ffaaff"},
	"CJ":                    {"CJ", "#ffaaff"},
	"CN":                    {"CN", "#ffaaff"},
	"D":                     {"D", "#ffaaff"},
	"DA":                    {"DA", "#ffaaff"},
	"DAB":                   {"DAB", "#ffaaff"},
	"DAV":                   {"DAV", "#ffaaff"},
	"DAZ":                   {"DAZ", "#ffaaff"},
	"DB":                    {"DB", "#ffaaff"},
	"DBV":                   {"DBV", "#ffaaff"},
	"DBZ":                   {"DBZ", "#ffaaff"},
	"DC":                    {"DC", "#ffaaff"},
	"DCV":                   {"DCV", "#ffaaff"},
	"DQ":                    {"DQ", "#ffaaff"},
	"F":                     {"F", "#ffaaff"},
	"F_WhiteSuperGiant":     {"F++", "#ffaaff"},
	"G":                     {"G", "#ffaaff"},
	"G_WhiteSuperGiant":     {"G++", "#ffaaff"},
	"H":                     {"H", "#ffaaff"},
	"K":                     {"K", "#ffaaff"},
	"K_OrangeGiant":         {"K+", "#ffaaff"},
	"L":                     {"L", "#ffaaff"},
	"M":                     {"M", "#ffaaff"},
	"M_RedGiant":            {"M+", "#ffaaff"},
	"M_RedSuperGiant":       {"M++", "#ffaaff"},
	"MS":                    {"MS", "#ffaaff"},
	"N":                     {"N", "#ffaaff"},
	"O":                     {"O", "#ffaaff"},
	"S":                     {"S", "#ffaaff"},
	"SupermassiveBlackHole": {"H++", "#ffaaff"},
	"T":                     {"T", "#ffaaff"},
	"TTS":                   {"TTS", "#ffaaff"},
	"W":                     {"W", "#ffaaff"},
	"WC":                    {"WC", "#ffaaff"},
	"WN":                    {"WN", "#ffaaff"},
	"WNC":                   {"WNC", "#ffaaff"},
	"WO":                    {"WO", "#ffaaff"},
	"Y":                     {"Y", "#ffaaff"},
}
