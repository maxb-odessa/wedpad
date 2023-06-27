package events

import (
	"errors"
	"fmt"
	"strings"
	"wedpad/internal/sound"
	"wedpad/internal/utils"

	"github.com/maxb-odessa/sconf"
)

type CurrentSystemT struct {
	name     string
	sector   string
	mainStar struct {
		id    int
		name  string
		typ   string
		color string
	}

	// systemwide, AutoScan, FSS*
	stars         map[int]*ScanT
	starsByName   map[string]*ScanT
	planets       map[int]*ScanT
	planetsByName map[string]*ScanT
	baryCentres   map[int]*ScanBaryCentreT
	signals       map[string]*FSSSignalDiscoveredT

	// bodywide, DSS, SAA
	planetSignalsCount map[int]*FSSBodySignalsT
	planetSignalsFound map[int]*SAASignalsFoundT

	bios *BiosT

	alert *AlertT
	sound *sound.SoundT
}

var loadedData map[string][]byte

func (cs *CurrentSystemT) Init() error {

	loadedData = make(map[string][]byte)

	// load all data: bios, geos, etc...
	if err := utils.LoadDir(loadedData, sconf.StrDef("resources", "data", "data"), ".json", 100_000, 32); err != nil {
		return err
	}

	// bios data loaded - init its handler
	if bioData, ok := loadedData["bios"]; ok {
		cs.bios = new(BiosT)
		if err := cs.bios.Init(bioData); err != nil {
			return err
		}
	} else {
		return errors.New("No BIO data loaded (missing 'bios.json'?)")
	}

	// load and init sounds
	cs.sound = new(sound.SoundT)
	if err := cs.sound.Init(); err != nil {
		return err
	}

	cs.alert = new(AlertT)
	if err := cs.alert.Init(); err != nil {
		return err
	}

	cs.Reset()

	return nil
}

func (cs *CurrentSystemT) Play(what string) {
	cs.sound.Play(what)
}

func (cs *CurrentSystemT) String() string {
	return cs.name
}

func (cs *CurrentSystemT) AddStar(s *ScanT) {
	cs.stars[s.BodyID] = s
	cs.starsByName[s.BodyName] = s
}

func (cs *CurrentSystemT) AddPlanet(p *ScanT) {
	cs.planets[p.BodyID] = p
	cs.planetsByName[p.BodyName] = p
}

func (cs *CurrentSystemT) AddSignal(s *FSSSignalDiscoveredT) {
	cs.signals[s.SignalName] = s
}

func (cs *CurrentSystemT) AddPlanetSignalsCount(s *FSSBodySignalsT) {
	cs.planetSignalsCount[s.BodyID] = s
}

func (cs *CurrentSystemT) AddPlanetSignalsFound(s *SAASignalsFoundT) {
	cs.planetSignalsFound[s.BodyID] = s
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

func (cs *CurrentSystemT) SetMainStarName(n string) {
	cs.mainStar.name = n
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

func (cs *CurrentSystemT) MainStarName() string {
	return cs.mainStar.name
}

func (cs *CurrentSystemT) Stars() map[int]*ScanT {
	return cs.stars
}

func (cs *CurrentSystemT) StarsByName() map[string]*ScanT {
	return cs.starsByName
}

func (cs *CurrentSystemT) BaryCentres() map[int]*ScanBaryCentreT {
	return cs.baryCentres
}

func (cs *CurrentSystemT) Planets() map[int]*ScanT {
	return cs.planets
}

func (cs *CurrentSystemT) PlanetsByName() map[string]*ScanT {
	return cs.planetsByName
}

func (cs *CurrentSystemT) Signals() map[string]*FSSSignalDiscoveredT {
	return cs.signals
}

func (cs *CurrentSystemT) PlanetSignalsCount() map[int]*FSSBodySignalsT {
	return cs.planetSignalsCount
}

func (cs *CurrentSystemT) PlanetSignalsFound() map[int]*SAASignalsFoundT {
	return cs.planetSignalsFound
}

func (cs *CurrentSystemT) Reset() {

	cs.name = "(somewhere in space)"
	cs.sector = "?"
	cs.mainStar = struct {
		id               int
		name, typ, color string
	}{0, "", "?", "#FFFFFF"}

	// systemwide
	cs.stars = make(map[int]*ScanT)
	cs.starsByName = make(map[string]*ScanT)
	cs.planets = make(map[int]*ScanT)
	cs.planetsByName = make(map[string]*ScanT)
	cs.baryCentres = make(map[int]*ScanBaryCentreT)
	cs.signals = make(map[string]*FSSSignalDiscoveredT, 0)

	// planet related
	cs.planetSignalsCount = make(map[int]*FSSBodySignalsT)
	cs.planetSignalsFound = make(map[int]*SAASignalsFoundT)
}

// in meters
const (
	SOLAR_RADIUS = 696_340_000.0
	EARTH_RADIUS = 6_371.0 * 1000.0
	LIGHT_SECOND = 299_792.0 * 1000.0
)

// in seconds
const (
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

func (cs *CurrentSystemT) BodyName(bname string) string {
	sysname := cs.Name()
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
	"B":                     {"B", "#E5E5FF"},
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
	"Y":                     {"Y", "#EE3090"},
	// special: undefined star type
	"?": {"&#10037;", "#A0A0A0"},
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
		return "#DD40A0" // T
	case temp >= 256:
		return "#DD3090" // Y
	}
	return "#DD3070" // too cold?
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
	"Gas giant with ammonia based life": {"GgABL", "#F5DEB3"},
	"Gas giant with water based life":   {"GgWBL", "#DEB887"},
	"Helium gas giant":                  {"HeGg", "#E2C49C"},
	"Helium rich gas giant":             {"HeRGg", "#DEA8C"},
	"High metal content body":           {"HMC", "#778899"},
	"Icy body":                          {"Icy", "#F8F8FF"},
	"Metal rich body":                   {"Metal", "#B0C4DE"},
	"Rocky body":                        {"Rocky", "#DEBA87"},
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

func small(num int) string {
	return fmt.Sprintf(`<span style="font-size: 0.7em;">%d</span>`, num)
}

var PlanetAtmospheres = map[string]TypeColorPair{
	"Ammonia":           {"NH" + small(3), "#B54A4A"},
	"AmmoniaOxygen":     {"NH" + small(3) + "+0" + small(2), "#BC8F8F"},
	"AmmoniaRich":       {"NH" + small(3) + "+", "#A52A2A"},
	"Argon":             {"Ar", "#3CB371"},
	"ArgonRich":         {"Ar+", "#3CB371"},
	"CarbonDioxide":     {"CO" + small(2), "#FF8C00"},
	"CarbonDioxideRich": {"CO" + small(2) + "+", "#FF8C00"},
	"EarthLike":         {"N" + small(2) + "+O" + small(2), "#ADD8E6"},
	"Helium":            {"He", "#DC143C"},
	"MetallicVapour":    {"Metal", "#D3D3D3"},
	"Methane":           {"CH" + small(4), "#FFDAB9"},
	"MethaneRich":       {"CH" + small(4) + "+", "#FFDAB9"},
	"Neon":              {"Ne", "#FF4500"},
	"NeonRich":          {"Ne+", "#FF4500"},
	"Nitrogen":          {"N" + small(2), "#87CEFA"},
	"None":              {"", "#000000"},
	"Oxygen":            {"O" + small(2), "#1E90FF"},
	"SilicateVapour":    {"SiO" + small(4) + "+", "#D3D3D3"},
	"SulphurDioxide":    {"SO" + small(2), "#CCCC00"},
	"Water":             {"H" + small(2) + "O", "#00BFFF"},
	"WaterRich":         {"H" + small(2) + "O+", "#00BFFF"},
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
