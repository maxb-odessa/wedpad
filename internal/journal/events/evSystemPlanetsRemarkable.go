package events

import (
	"fmt"
	"strings"

	"github.com/danwakefield/fnmatch"
	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

func (cs *CurrentSystemT) IsRemarkableBody(id int) bool {

	name := cs.Planets()[id].BodyName

	// TODO mark as remarkable by close bodies, shepherd moon, close rings, etc
	// tap on planet name for detailed info popup (including bios)

	if cs.wantBGGHO(id) {
		slog.Debug(5, "Remarkable body '%s': wantBGGHO", name)
		return true
	}

	if cs.wantRings(id) {
		slog.Debug(5, "Remarkable body '%s': wantRings", name)
		return true
	}

	if cs.wantAtmosphere(id) {
		slog.Debug(5, "Remarkable body '%s': wantAtmosphere", name)
		return true
	}

	if cs.wantDMTL(id) {
		slog.Debug(5, "Remarkable body '%s': wantDMTL", name)
		return true
	}

	if cs.wantGravity(id) {
		slog.Debug(5, "Remarkable body '%s': wantGravity", name)
		return true
	}

	if cs.wantBodyClass(id) {
		slog.Debug(5, "Remarkable body '%s': wantBodyClass", name)
		return true
	}

	slog.Debug(9, "Remarkable body '%s': NOT REMARKABLE", name)

	return false
}

func (cs *CurrentSystemT) wantBGGHO(id int) bool {

	wantBio := int(sconf.Int32Def("criteria", "min bio signals", 0))
	wantGeo := int(sconf.Int32Def("criteria", "min geo signals", 0))
	wantGuard := int(sconf.Int32Def("criteria", "min guardian signals", 0))
	wantHuman := int(sconf.Int32Def("criteria", "min human signals", 0))
	wantOther := int(sconf.Int32Def("criteria", "min other signals", 0))

	// no signals detected but at least one is wanted
	sigs, ok := cs.PlanetSignalsCount()[id]
	if !ok {
		if wantBio+wantGeo+wantGuard+wantHuman+wantOther > 0 {
			return false
		} else {
			return true
		}
	}

	for _, sig := range sigs.Signals {
		switch sig.Type {
		case "$SAA_SignalType_Biological;":
			if sig.Count >= wantBio {
				return true
			}
		case "$SAA_SignalType_Geological;":
			if sig.Count >= wantGeo {
				return true
			}
		case "$SAA_SignalType_Guardian;":
			if sig.Count >= wantGuard {
				return true
			}
		case "$SAA_SignalType_Human;":
			if sig.Count >= wantHuman {
				return true
			}
		case "$SAA_SignalType_Other;":
			if sig.Count >= wantOther {
				return true
			}
		}

	}

	return false
}

func (cs *CurrentSystemT) wantRings(id int) bool {

	rNum, rRad := CalcRings(cs.Planets()[id])
	rRadLs := float32(rRad / LIGHT_SECOND)

	wantHighRadLs := sconf.Float32Def("criteria", "high ring radius", 0.0)
	wantLowRadLs := sconf.Float32Def("criteria", "low ring radius", 0.0)

	if rNum == 0 {
		if wantLowRadLs > 0.0 || wantHighRadLs > 0.0 {
			return false
		} else {
			return true
		}
	}

	if rRadLs <= wantLowRadLs || rRadLs >= wantHighRadLs {
		return true
	}

	return false
}

func (cs *CurrentSystemT) wantAtmosphere(id int) bool {

	body := cs.Planets()[id]

	wantAtmos := sconf.StrDef("criteria", "atmospheres", "*")
	landableOnly := sconf.BoolDef("criteria", "atmospheres landable only", false)
	found := false

	for _, atmo := range strings.Split(wantAtmos, ",") {
		if fnmatch.Match(strings.TrimSpace(atmo), body.Atmosphere, fnmatch.FNM_IGNORECASE) {
			found = true
			break
		}
	}

	if found {
		if landableOnly && !body.Landable {
			found = false
		}
	}

	return found
}

func (cs *CurrentSystemT) wantDMTL(id int) bool {

	body := cs.Planets()[id]

	wantDiscovered := sconf.BoolDef("criteria", "discovered", false)
	wantMapped := sconf.BoolDef("criteria", "mapped", false)
	wantTerraformable := sconf.BoolDef("criteria", "terraformable", false)
	wantLandable := sconf.BoolDef("criteria", "landable", false)

	if wantDiscovered && body.WasDiscovered ||
		wantMapped && body.WasMapped ||
		wantTerraformable && body.TerraformState == "Terraformable" ||
		wantLandable && body.Landable {
		return true
	}

	return false
}

func (cs *CurrentSystemT) wantGravity(id int) bool {

	body := cs.Planets()[id]

	wantHighGravity := sconf.Float32Def("criteria", "high gravity", 0.0)
	wantLowGravity := sconf.Float32Def("criteria", "low gravity", 0.0)
	landableOnly := sconf.BoolDef("criteria", "gravity landable only", false)

	grav := float32(body.SurfaceGravity / 10.0) // yes, 10.0 is corrent here (as FDEV why)

	if grav <= wantLowGravity || grav >= wantHighGravity {
		if landableOnly && !body.Landable {
			return false
		} else {
			return true
		}
	}

	return false
}

func (cs *CurrentSystemT) wantBodyClass(id int) bool {

	body := cs.Planets()[id]

	wantClass := sconf.StrDef("criteria", "body class", "*")

	for _, class := range strings.Split(wantClass, ",") {
		if fnmatch.Match(strings.TrimSpace(class), body.PlanetClass, fnmatch.FNM_IGNORECASE) {
			return true
		}
	}

	return false
}

func (cs *CurrentSystemT) notesOnBody(id int) []string {

	notes := make([]string, 0)

	// close bodies
	if n := closeBodies(cs, id); n != "" {
		notes = append(notes, n)
	}

	// close/intersecting rings

	// shepherd moon

	// hot planet (close to star)

	// fast spinning body

	// fast orbiting body

	// high inclination body

	slog.Debug(9, "Notes on body: %+v", notes)

	return notes
}

func closeBodies(cs *CurrentSystemT, id int) string {

	planets := cs.Planets()
	body := planets[id]

	if parent := findParentBody(planets, id); parent != nil {

		bodyDistRatio := body.SemiMajorAxis / body.Radius
		parentDistRatio := body.SemiMajorAxis / parent.Radius

		if bodyDistRatio <= 5.0 || parentDistRatio <= 5.0 {
			return fmt.Sprintf("Close bodies '%s' and '%s', SMA/Rad: %.2f and %.2f", cs.BodyName(body.BodyName), cs.BodyName(parent.BodyName), bodyDistRatio, parentDistRatio)
		}

	}

	return ""
}

func findParentBody(planets map[int]*ScanT, id int) *ScanT {

	body := planets[id]

	for _, parent := range body.Parents {
		if parent.Planet > 0 {
			return planets[parent.Planet]
		}
	}

	return nil
}
