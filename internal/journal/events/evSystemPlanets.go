package events

import (
	"sort"
	"strconv"
	"wedpad/internal/msg"

	"github.com/maxb-odessa/slog"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func (cs *CurrentSystemT) ShowPlanets() {

	// those are for BODIES view port
	bodies := make([]map[string]interface{}, 0)

	planets := cs.Planets()

	var keys []int

	for k, _ := range planets {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, id := range keys {

		b := planets[id]

		if !cs.remarkable("planet", b) {
			continue
		}

		body := make(map[string]interface{})

		body["Name"] = BodyName(b.BodyName, cs.Name())
		body["DistanceLs"] = Num(b.DistanceFromArrivalLs)
		body["Type"] = PlanetTypeColor(b.PlanetClass)
		body["DMTL"] = composeDMTL(b.WasDiscovered, b.WasMapped, b.TerraformState, b.Landable)
		body["RadiusE"] = Num(b.Radius / EARTH_RADIUS)
		body["MassE"] = Num(b.MassEm)
		body["TemperatureK"] = Num(b.SurfaceTemperature)
		body["GravityG"] = Num(b.SurfaceGravity / 10) // this 10 is correct!
		body["RingsNum"], body["RingsRadLs"] = CalcRings(b)
		body["Atmosphere"] = PlanetAtmosphereTypeColor(b.AtmosphereType)
		body["Signals"] = cs.composeSignals(id)

		bodies = append(bodies, body)

	}

	mb := &msg.Message{
		Target: msg.TARGET_BODIES,
		Type:   msg.TYPE_VIEW,
		Action: msg.ACTION_REPLACE,
		Data:   bodies,
	}

	mb.Send()

	// those are for LOG view port
	for _, b := range cs.planets {

		if !cs.remarkable("body", b) {
			continue
		}

	}

}

func composeDMTL(d bool, m bool, tf string, l bool) []TypeColorPair {
	dmtl := make([]TypeColorPair, 4)

	if d == true {
		dmtl[0] = TypeColorPair{Type: "D", Color: "yellow"}
	} else {
		dmtl[0] = TypeColorPair{Type: "-", Color: "gray"}
	}

	if m == true {
		dmtl[1] = TypeColorPair{Type: "M", Color: "yellow"}
	} else {
		dmtl[1] = TypeColorPair{Type: "-", Color: "gray"}
	}

	if tf == "Terraformable" {
		dmtl[2] = TypeColorPair{Type: "T", Color: "lime"}
	} else {
		dmtl[2] = TypeColorPair{Type: "-", Color: "gray"}
	}

	if l == true {
		dmtl[3] = TypeColorPair{Type: "L", Color: "cyan"}
	} else {
		dmtl[3] = TypeColorPair{Type: "-", Color: "gray"}
	}

	return dmtl
}

func (cs *CurrentSystemT) composeSignals(id int) (signals []TypeColorPair) {
	signals = make([]TypeColorPair, 3)

	// init with defaults
	for i, _ := range signals {
		signals[i] = TypeColorPair{Type: "-", Color: "gray"}
	}

	// no signals detected
	sigs, ok := cs.PlanetSignals()[id]
	if !ok {
		return
	}

	var bio, geo, other int
	for _, sig := range sigs.Signals {
		switch sig.Type {
		case "$SAA_SignalType_Biological":
			bio++
		case "$SAA_SignalType_Geological":
			geo++
		default:
			other++
		}
	}

	if bio > 0 {
		signals[0] = TypeColorPair{Type: strconv.Itoa(bio), Color: "green"}
	}

	if geo > 0 {
		signals[1] = TypeColorPair{Type: strconv.Itoa(geo), Color: "brown"}
	}

	if other > 0 {
		signals[2] = TypeColorPair{Type: strconv.Itoa(bio), Color: "blue"}
	}

	return
}

type luaData struct {
	ev  *ScanT
	cs  *CurrentSystemT
	res bool
}

func (d *luaData) SetResult(r bool) {
	d.res = r
}

func (cs *CurrentSystemT) remarkable(what string, ev *ScanT) bool {

	scriptName := "remarkable-" + what

	script, err := cs.FindLua(scriptName)
	if err != nil {
		slog.Err("LUA: Failed to call '%s.lua': %s", scriptName, err)
		return false
	}

	l := lua.NewState()
	defer l.Close()

	ld := &luaData{
		ev:  ev,
		cs:  cs,
		res: false,
	}

	l.SetGlobal("data", luar.New(l, ld))
	if err := l.DoString(script); err != nil {
		slog.Err("LUA: Failed to call '%s.lua': %s", scriptName, err)
		return false
	}

	slog.Debug(9, "LUA: '%s' returned '%+v'", scriptName, ld.res)

	return ld.res
}
