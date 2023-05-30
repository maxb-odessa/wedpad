package events

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"wedpad/internal/msg"
)

func (cs *CurrentSystemT) ShowPlanets(final bool) {

	bodies := make([]map[string]interface{}, 0)

	planets := cs.Planets()

	var keys []int

	for k, _ := range planets {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	bodiesCnt := 0
	haveNotes := false

	for _, id := range keys {

		body := make(map[string]interface{})

		b := planets[id]

		if cs.IsRemarkableBody(id) {

			body["Name"] = cs.BodyName(b.BodyName)
			body["DistanceLs"] = Num(b.DistanceFromArrivalLs)
			body["Type"] = PlanetTypeColor(b.PlanetClass)
			body["DMTL"] = composeDMTL(b.WasDiscovered, b.WasMapped, b.TerraformState, b.Landable)
			body["RadiusE"] = Num(b.Radius / EARTH_RADIUS)
			body["MassE"] = Num(b.MassEm)
			body["TemperatureK"] = Num(b.SurfaceTemperature)
			body["GravityG"] = Num(b.SurfaceGravity / 10) // this 10 is correct!
			rn, rr := CalcRings(b)
			if rn > 0 {
				body["Rings"] = fmt.Sprintf("%d/%.1f", rn, rr/LIGHT_SECOND)
			}
			body["Atmosphere"] = PlanetAtmosphereTypeColor(b.AtmosphereType)
			body["Signals"] = cs.composeBGGHO(id)

			bodies = append(bodies, body)
			bodiesCnt++
		}

		if final {
			if notes := cs.notesOnBody(id); len(notes) > 0 {
				body["Notes"] = strings.Join(notes, "<br>")
				body["NotesName"] = cs.BodyName(b.BodyName)
				haveNotes = true
				bodies = append(bodies, body)
			}
		}

	}

	if bodiesCnt > 0 || haveNotes {

		m := &msg.Message{
			Target: msg.TARGET_BODIES,
			Type:   msg.TYPE_VIEW,
			Action: msg.ACTION_REPLACE,
			Data:   bodies,
		}
		m.Send()

		buttonText := fmt.Sprintf("%d", bodiesCnt)
		if haveNotes {
			buttonText += " (!)"
		}
		m = &msg.Message{
			Target: msg.TARGET_BODIES,
			Type:   msg.TYPE_BUTTON,
			Action: msg.ACTION_REPLACE,
			Data:   buttonText,
		}
		m.Send()

		m = &msg.Message{
			Target: msg.TARGET_BODIES,
			Type:   msg.TYPE_VIEW,
			Action: msg.ACTION_ATTENTION,
			Data:   "",
		}
		m.Send()
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
	sigs, ok := cs.PlanetSignalsCount()[id]
	if !ok {
		return
	}

	var bio, geo, guard, human, other int
	for _, sig := range sigs.Signals {
		switch sig.Type {
		case "$SAA_SignalType_Biological;":
			bio = sig.Count
		case "$SAA_SignalType_Geological;":
			geo = sig.Count
		case "$SAA_SignalType_Guardian;":
			guard = sig.Count
		case "$SAA_SignalType_Human;":
			human = sig.Count
		case "$SAA_SignalType_Other;":
			other = sig.Count
		}
	}

	if bio > 0 {
		signals[0] = TypeColorPair{Type: strconv.Itoa(bio), Color: "lime"}
	}

	if geo > 0 {
		signals[1] = TypeColorPair{Type: strconv.Itoa(geo), Color: "sandybrown"}
	}

	if guard > 0 {
		signals[2] = TypeColorPair{Type: strconv.Itoa(guard), Color: "yellow"}
	}

	if human > 0 {
		signals[3] = TypeColorPair{Type: strconv.Itoa(human), Color: "dodgerblue"}
	}

	if other > 0 {
		signals[4] = TypeColorPair{Type: strconv.Itoa(other), Color: "red"}
	}

	return
}
