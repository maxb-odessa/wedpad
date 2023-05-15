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
	signals       map[int]*SAASignalsFoundT
	planetSignals map[int]*FSSBodySignalsT
}

func (cs *CurrentSystemT) Init() error {
	cs.Clean("all")
	return nil
}

func (cs *CurrentSystemT) String() string {
	return cs.name
}

func (cs *CurrentSystemT) AddStar(s *ScanT) {
	cs.stars[s.BodyID] = s
}

func (cs *CurrentSystemT) AddPlanet(p *ScanT) {
	cs.planets[p.BodyID] = p
}

func (cs *CurrentSystemT) AddPlanetSignals(s *FSSBodySignalsT) {
	cs.planetSignals[s.BodyID] = s
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
	scp := StarTypeColor(t)
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

func (cs *CurrentSystemT) PlanetSignals() map[int]*FSSBodySignalsT {
	return cs.planetSignals
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
		cs.signals = make(map[int]*SAASignalsFoundT)
	case "planetSignals":
		cs.planetSignals = make(map[int]*FSSBodySignalsT)
	case "all":
		cs.stars = make(map[int]*ScanT)
		cs.planets = make(map[int]*ScanT)
		cs.baryCentres = make(map[int]*ScanBaryCentreT)
		cs.signals = make(map[int]*SAASignalsFoundT)
		cs.planetSignals = make(map[int]*FSSBodySignalsT)
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

func BodyName(bname, sysname string) string {
	if bname == sysname {
		return ""
	}
	return strings.TrimPrefix(bname, sysname+" ")
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

type TypeColorPair struct {
	Type, Color string
}

var StarTypes = map[string]TypeColorPair{
	"A":                     {"A", "#FFFFE0"},
	"A_BlueWhiteSuperGiant": {"++A", "#FFFFE0"},
	"AeBe":                  {"AeBe", "#FFFF80"},
	"B":                     {"B", "#E0E0FF"},
	"B_BlueWhiteSuperGiant": {"++B", "#E0E0FF"},
	"C":                     {"C", "#DDDD60"},
	"CJ":                    {"CJ", "#DDDD60"},
	"CN":                    {"CN", "#DDDD60"},
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
	"S":                     {"S", "#DDDD70"},
	"SupermassiveBlackHole": {"++H", "#808080"},
	"T":                     {"T", "#EE2080"},
	"TTS":                   {"TTS", ""},
	"W":                     {"W", "#E0E0E0"},
	"WC":                    {"WC", "#E0E0E0"},
	"WN":                    {"WN", "#E0E0E0"},
	"WNC":                   {"WNC", "#E0E0E0"},
	"WO":                    {"WO", "#E0E0E0"},
	"Y":                     {"Y", "#DD2080"},
	// special: undefined star type
	"?": {"(UNK)", "#A0A0A0"},
}

func GuessColorByTemp(temp float64) string {
	switch {
	case temp >= 50_000:
		return "#AAAAFF" // N
	case temp >= 33_000:
		return "#EEEEFF" // O
	case temp >= 10_000:
		return "#E0E0FF" // B
	case temp >= 6_000:
		return "#FFFFB0" // F
	case temp >= 5_200:
		return "#EEEE30" // G
	case temp >= 3_700:
		return "#FF9020" // K
	case temp >= 2_000:
		return "#FF3030" // M
	case temp >= 1_300:
		return "#EE2080" // L
	case temp >= 999:
		return "#CC2080" // T
	case temp >= 256:
		return "#AA2080" // Y
	}
	return "#aabbcc" // dunno
}

func StarTypeColor(st string) TypeColorPair {
	if tc, ok := StarTypes[st]; ok {
		return tc
	}
	// star type is new to us!
	return TypeColorPair{
		Type:  st + "(1fixme!)",
		Color: "#FFFF00",
	}
}

var PlanetTypes = map[string]TypeColorPair{
	"Ammonia world":                     {"AW", "#F4A460"},
	"Earthlike body":                    {"ELW", "#00FF7F"},
	"Gas giant with ammonia based life": {"GgAL", "#F5DEB3"},
	"Gas giant with water based life":   {"GgWL", "#DEB887"},
	"Helium gas giant":                  {"HeGg", "#D2B48C"},
	"Helium rich gas giant":             {"HeRGg", "#D2B48C"},
	"High metal content body":           {"HMC", "#778899"},
	"Icy body":                          {"Icy", "#F8F8FF"},
	"Metal rich body":                   {"Metal", "#B0C4DE"},
	"Rocky body":                        {"Rocky", "#DEB887"},
	"Rocky ice body":                    {"RoIce", "#FFF8DC"},
	"Sudarsky class I gas giant":        {"GgI", "#DEB887"},
	"Sudarsky class II gas giant":       {"GgII", "#F5DEB3"},
	"Sudarsky class III gas giant":      {"GgIII", "#6495ED"},
	"Sudarsky class IV gas giant":       {"GgIV", "#CD5C5C"},
	"Sudarsky class V gas giant":        {"GgV", "#DC143C"},
	"Water giant":                       {"Wg", "#808000"},
	"Water world":                       {"WW", "#00BFFF"},
}

func PlanetTypeColor(pt string) TypeColorPair {
	if tc, ok := PlanetTypes[pt]; ok {
		return tc
	}
	// body type is new to us!
	return TypeColorPair{
		Type:  pt + "(2fixme!)",
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

var PlanetAtmospheres = map[string]TypeColorPair{
	"Ammonia":           {"NH3", "#A52A2A"},
	"AmmoniaOxygen":     {"NH3+02", "#BC8F8F"},
	"AmmoniaRich":       {"NH3+", "#A52A2A"},
	"Argon":             {"Ar", "#3CB371"},
	"ArgonRich":         {"Ar+", "#3CB371"},
	"CarbonDioxide":     {"CO2", "#FF8C00"},
	"CarbonDioxideRich": {"CO2+", "#FF8C00"},
	"EarthLike":         {"N2+O2", "#ADD8E6"},
	"Helium":            {"He", "#DC143C"},
	"MetallicVapour":    {"Metals", "#D3D3D3"},
	"Methane":           {"CH4", "#FFDAB9"},
	"MethaneRich":       {"CH4+", "#FFDAB9"},
	"Neon":              {"Ne", "#FF4500"},
	"NeonRich":          {"Ne+", "#FF4500"},
	"Nitrogen":          {"N2", "#87CEFA"},
	"None":              {"", "#000000"},
	"Oxygen":            {"O2", "#1E90FF"},
	"SilicateVapour":    {"Si", "#D3D3D3"},
	"SulphurDioxide":    {"SO2", "#CCCC00"},
	"Water":             {"H2O", "#00BFFF"},
	"WaterRich":         {"H2O+", "#00BFFF"},
	"":                  {"", "#000000"},
}

func PlanetAtmosphereTypeColor(pa string) TypeColorPair {
	if tc, ok := PlanetAtmospheres[pa]; ok {
		return tc
	}
	// atmo is new to us!
	return TypeColorPair{
		Type:  pa + "(fixme!)",
		Color: "#FFFF00",
	}
}
