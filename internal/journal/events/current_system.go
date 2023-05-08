package events

import (
	"fmt"
	"strings"
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

func (cs *CurrentSystemT) String() string {
	return cs.Name
}

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
		s = fmt.Sprintf("%3.2f", val)
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

type StarTypeColorPair struct {
	Type, Color string
}

var StarTypes = map[string]StarTypeColorPair{
	"A":                     {"A", "#FFFFE0"},
	"A_BlueWhiteSuperGiant": {"++A", "#FFFFE0"},
	"AeBe":                  {"AeBe", "#FFFF80"},
	"B":                     {"B", "#E0E0FF"},
	"B_BlueWhiteSuperGiant": {"++B", "#E0E0FF"},
	"C":                     {"C", "#DDDD70"},
	"CJ":                    {"CJ", "#DDDD70"},
	"CN":                    {"CN", "#DDDD70"},
	"D":                     {"D", "#EEEEFF"},
	"DA":                    {"DA", "#EEEEFF"},
	"DAB":                   {"DAB", "#EEEEFF"},
	"DAV":                   {"DAV", "#EEEEFF"},
	"DAZ":                   {"DAZ", "#EEEEFF"},
	"DB":                    {"DB", "#EEEEFF"},
	"DBV":                   {"DBV", "#EEEEFF"},
	"DBZ":                   {"DBZ", "#EEEEFF"},
	"DC":                    {"DC", "#EEEEFF"},
	"DCV":                   {"DCV", "#EEEEFF"},
	"DQ":                    {"DQ", "#EEEEFF"},
	"F":                     {"F", "#FFFFA0"},
	"F_WhiteSuperGiant":     {"++F", "#FFFFA0"},
	"G":                     {"G", "#EEEE30"},
	"G_WhiteSuperGiant":     {"++G", "#EEEE30"},
	"H":                     {"H", "#808080"},
	"K":                     {"K", "#FF9020"},
	"K_OrangeGiant":         {"+K", "#FF9020"},
	"L":                     {"L", "#FF2080"},
	"M":                     {"M", "#FF3030"},
	"M_RedGiant":            {"+M", "#FF3030"},
	"M_RedSuperGiant":       {"++M", "#FF3030"},
	"MS":                    {"MS", "#FF7010"},
	"N":                     {"N", "#AAAAFF"},
	"O":                     {"O", "#FFFFFF"},
	"S":                     {"S", "#CCCC70"},
	"SupermassiveBlackHole": {"++H", "#808080"},
	"T":                     {"T", "#FF2080"},
	"TTS":                   {"TTS", "#FF7070"},
	"W":                     {"W", "#E0E0E0"},
	"WC":                    {"WC", "#E0E0E0"},
	"WN":                    {"WN", "#E0E0E0"},
	"WNC":                   {"WNC", "#E0E0E0"},
	"WO":                    {"WO", "#E0E0E0"},
	"Y":                     {"Y", "#FF2080"},
}

func StarTypeColor(t string) StarTypeColorPair {
	if st, ok := StarTypes[t]; ok {
		return st
	}
	return StarTypeColorPair{
		Type:  t + "(?!)",
		Color: "#FFFF00",
	}
}
