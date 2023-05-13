package events

import (
	"sort"
	"strconv"
	"wedpad/internal/msg"
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

	bodiesCnt := 0
	for _, id := range keys {

		b := planets[id]

		if !cs.IsRemarkableBody(id) {
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
		rn, rr := CalcRings(b)
		if rn > 0 {
			body["RingsNum"] = rn
			body["RingsRadiusLs"] = Num(rr / LIGHT_SECOND)
		}
		body["Atmosphere"] = PlanetAtmosphereTypeColor(b.AtmosphereType)
		body["Signals"] = cs.composeBGGHO(id)

		bodies = append(bodies, body)

		bodiesCnt++
	}

	if bodiesCnt > 0 {

		m := &msg.Message{
			Target: msg.TARGET_BODIES,
			Type:   msg.TYPE_VIEW,
			Action: msg.ACTION_REPLACE,
			Data:   bodies,
		}
		m.Send()

		m = &msg.Message{
			Target: msg.TARGET_BODIES,
			Type:   msg.TYPE_BUTTON,
			Action: msg.ACTION_REPLACE,
			Data:   bodiesCnt,
		}
		m.Send()

		m = &msg.Message{
			Target: msg.TARGET_BODIES,
			Type:   msg.TYPE_BUTTON,
			Action: msg.ACTION_ATTENTION,
			Data:   "",
		}
		m.Send()
	}

	// those are for LOG view
	for _, id := range keys {

		if !cs.IsRemarkableBody(id) {
			continue
		}

	}

}

// DMTL = Discovered, Mapped, Terraformable, Landable
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

// BGGHO = Bio, Geo, Guardian, Human, Other
func (cs *CurrentSystemT) composeBGGHO(id int) (signals []TypeColorPair) {
	signals = make([]TypeColorPair, 5)

	// init with defaults
	for i, _ := range signals {
		signals[i] = TypeColorPair{Type: "-", Color: "gray"}
	}

	// no signals detected
	sigs, ok := cs.PlanetSignals()[id]
	if !ok {
		return
	}

	var bio, geo, guard, human, other int
	for _, sig := range sigs.Signals {
		switch sig.Type {
		case "$SAA_SignalType_Biological":
			bio++
		case "$SAA_SignalType_Geological":
			geo++
		case "$SAA_SignalType_Guardian":
			guard++
		case "$SAA_SignalType_Human":
			human++
		case "$SAA_SignalType_Other":
			other++
		}
	}

	if bio > 0 {
		signals[0] = TypeColorPair{Type: strconv.Itoa(bio), Color: "green"}
	}

	if geo > 0 {
		signals[1] = TypeColorPair{Type: strconv.Itoa(geo), Color: "brown"}
	}

	if guard > 0 {
		signals[2] = TypeColorPair{Type: strconv.Itoa(guard), Color: "yellow"}
	}

	if human > 0 {
		signals[3] = TypeColorPair{Type: strconv.Itoa(human), Color: "blue"}
	}

	if other > 0 {
		signals[4] = TypeColorPair{Type: strconv.Itoa(other), Color: "red"}
	}

	return
}
