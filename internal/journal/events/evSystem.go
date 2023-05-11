package events

import (
	"fmt"
	"strings"
)

type CurrentSystemT struct {
	name     string
	sector   string
	mainStar struct {
		id    int
		typ   string
		color string
	}
	stars         map[int]*ScanT
	planets       map[int]*ScanT
	baryCentres   map[int]*ScanBaryCentreT
	signals       map[int]*FSSSignalDiscoveredT
	planetSignals map[int]*SAASignalsFoundT
	luaScripts    map[string]string
}

func (cs *CurrentSystemT) Init() error {
	cs.Clean("all")
	return cs.LoadLua()
}

func (cs *CurrentSystemT) LoadLua() error {
	// TODO
	return nil
}

func (cs *CurrentSystemT) AddStar(s *ScanT) {
	cs.stars[s.BodyID] = s
}

func (cs *CurrentSystemT) AddPlanet(p *ScanT) {
	cs.planets[p.BodyID] = p
}

func (cs *CurrentSystemT) AddBaryCentre(b *ScanBaryCentreT) {
	cs.baryCentres[b.BodyID] = b
}

func (cs *CurrentSystemT) SetName(n string) {
	cs.name = n
}

func (cs *CurrentSystemT) SetMainStarID(id int) {
	cs.mainStar.id = id
}

func (cs *CurrentSystemT) Name() string {
	return cs.name
}

func (cs *CurrentSystemT) SetMainStarType(t string) {
	scp := cs.StarTypeColor(t)
	cs.mainStar.typ = scp.Type
	cs.mainStar.color = scp.Color
}

func (cs *CurrentSystemT) MainStarType() string {
	return cs.mainStar.typ
}

func (cs *CurrentSystemT) MainStarID() int {
	return cs.mainStar.id
}

func (cs *CurrentSystemT) Stars() map[int]*ScanT {
	return cs.stars
}

func (cs *CurrentSystemT) BaryCentres() map[int]*ScanBaryCentreT {
	return cs.baryCentres
}

func (cs *CurrentSystemT) Planets() map[int]*ScanT {
	return cs.planets
}

func (cs *CurrentSystemT) Clean(what string) {

	cs.name = "(somewhere in space)"
	cs.sector = "?"
	cs.mainStar = struct {
		id         int
		typ, color string
	}{0, "?", "#FFFFFF"}

	switch strings.ToLower(what) {
	case "stars":
		cs.stars = make(map[int]*ScanT)
	case "planets":
		cs.planets = make(map[int]*ScanT)
	case "barycentres":
		cs.baryCentres = make(map[int]*ScanBaryCentreT)
	case "signals":
		cs.signals = make(map[int]*FSSSignalDiscoveredT)
	case "planetSignals":
		cs.planetSignals = make(map[int]*SAASignalsFoundT)
	default:
		cs.stars = make(map[int]*ScanT)
		cs.planets = make(map[int]*ScanT)
		cs.baryCentres = make(map[int]*ScanBaryCentreT)
		cs.signals = make(map[int]*FSSSignalDiscoveredT)
		cs.planetSignals = make(map[int]*SAASignalsFoundT)
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
	"F":                     {"F", "#FFFFB0"},
	"F_WhiteSuperGiant":     {"++F", "#FFFFB0"},
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
	"O":                     {"O", "#EEEEFF"},
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
	// special: undefined star type
	"?": {"(UNK)", "#A0A0A0"},
}

func (cs *CurrentSystemT) StarTypeColor(t string) StarTypeColorPair {
	if st, ok := StarTypes[t]; ok {
		return st
	}
	// star type is new to us!
	return StarTypeColorPair{
		Type:  t + "(fixme!)",
		Color: "#FFFF00",
	}
}

/* not used atm
func StarTempColor(tempK float64) string {

	// black holes have no temp so make them dark gray
	if tempK == 0.0 {
		return "#101010"
	}

	r, g, b := temperature.ToRGB(uint16(tempK))
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}
*/
