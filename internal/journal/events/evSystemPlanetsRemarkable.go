package events

import (
	"fmt"
	"math"
	"strings"

	"github.com/danwakefield/fnmatch"
	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

func (cs *CurrentSystemT) IsRemarkableBody(id int) bool {

	name := cs.Planets()[id].BodyName

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

	minRadLs := sconf.Float32Def("criteria", "min ring radius", 0.0)

	if rNum == 0 {
		if minRadLs > 0.0 {
			return false
		} else {
			return true
		}
	}

	if rRadLs >= minRadLs {
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

	body := cs.Planets()[id]

	// close bodies or rings
	if note := cs.closeBodies(body); note != "" {
		notes = append(notes, note)
	}

	// shepherd moon
	if note := cs.shepherdMoon(body); note != "" {
		notes = append(notes, note)
	}

	// hot planet (close to star)
	if note := cs.hotPlanet(body); note != "" {
		notes = append(notes, note)
	}

	// fast spinning body
	if note := cs.fastSpinning(body); note != "" {
		notes = append(notes, note)
	}

	// high inclination body
	if note := cs.highInclination(body); note != "" {
		notes = append(notes, note)
	}

	// GG with high helium level
	if note := cs.highHeliumLevel(body); note != "" {
		notes = append(notes, note)
	}

	slog.Debug(9, "Notes on body: %+v", notes)

	return notes
}

func findParentBarycentre(barycentres map[int]*ScanBaryCentreT, body *ScanT) *ScanBaryCentreT {

	if len(body.Parents) > 0 {
		if body.Parents[0].Null > 0 {
			return barycentres[body.Parents[0].Null]
		}
	}

	return nil
}

func findParentBody(planets map[int]*ScanT, body *ScanT) *ScanT {

	if len(body.Parents) > 0 {
		if body.Parents[0].Planet > 0 {
			return planets[body.Parents[0].Planet]
		}
	}

	return nil
}

func findParentStar(stars map[int]*ScanT, body *ScanT) *ScanT {

	if len(body.Parents) > 0 {
		if body.Parents[0].Star >= 0 {
			return stars[body.Parents[0].Star]
		}
	}

	return nil
}

func findBodyByBaryCentreID(planets map[int]*ScanT, baryCentreID int, skipBodyID int) *ScanT {

	for id, p := range planets {
		if id == skipBodyID {
			continue
		}
		if len(p.Parents) > 0 {
			if p.Parents[0].Null == baryCentreID {
				return p
			}
		}
	}

	return nil
}

func (cs *CurrentSystemT) closeBodies(body *ScanT) string {

	ratioRequired := float64(sconf.Float32Def("criteria", "close bodies ratio", 3.0))

	// check over parent barycentre
	if pBary := findParentBarycentre(cs.BaryCentres(), body); pBary != nil {

		if pBody := findBodyByBaryCentreID(cs.Planets(), pBary.BodyID, body.BodyID); pBody != nil {

			bodyRad := body.Radius
			pBodyRad := pBody.Radius
			byRings := ""

			if rn, rr := CalcRings(body); rn > 0 {
				bodyRad += rr
				byRings = " (rings)"
			}

			if rn, rr := CalcRings(pBody); rn > 0 {
				pBodyRad += rr
				byRings = " (rings)"
			}

			bodyDistRatio := body.SemiMajorAxis / bodyRad
			parentDistRatio := pBody.SemiMajorAxis / pBodyRad

			if (bodyDistRatio < ratioRequired || parentDistRatio < ratioRequired) &&
				math.Abs(bodyDistRatio+parentDistRatio) < ratioRequired*3 {
				return fmt.Sprintf("Close orbiting bodes%s: to '%s', SMA/Rad: %.2f",
					byRings,
					cs.BodyName(pBody.BodyName),
					parentDistRatio,
				)
			}

		}

		// check over parent body
	} else if pPlanet := findParentBody(cs.Planets(), body); pPlanet != nil {

		bodyRad := body.Radius
		pPlanetRad := pPlanet.Radius
		byRings := ""

		if rn, rr := CalcRings(body); rn > 0 {
			bodyRad += rr
			byRings = " (rings)"
		}

		if rn, rr := CalcRings(pPlanet); rn > 0 {
			pPlanetRad += rr
			byRings = " (rings)"
		}

		bodyDistRatio := body.SemiMajorAxis / bodyRad
		parentDistRatio := body.SemiMajorAxis / pPlanetRad

		if (bodyDistRatio < ratioRequired || parentDistRatio < ratioRequired) &&
			math.Abs(bodyDistRatio+parentDistRatio) < ratioRequired*3 {
			return fmt.Sprintf("Close orbiting body%s: to '%s', SMA/Rad: %.2f (%.2f)",
				byRings,
				cs.BodyName(pPlanet.BodyName),
				bodyDistRatio,
				parentDistRatio,
			)
		}
	}

	return ""
}

func (cs *CurrentSystemT) shepherdMoon(body *ScanT) string {

	if parent := findParentBody(cs.Planets(), body); parent != nil {
		if rn, rr := CalcRings(parent); rn > 0 {
			if rr > body.SemiMajorAxis {
				return fmt.Sprintf("Shepherd Moon: for '%s'", cs.BodyName(parent.BodyName))
			}
		}
	}

	return ""
}

func (cs *CurrentSystemT) hotPlanet(body *ScanT) string {

	if parent := findParentStar(cs.Stars(), body); parent != nil {

		smaRadRatio := body.SemiMajorAxis / parent.Radius

		if smaRadRatio < float64(sconf.Float32Def("criteria", "hot planet ratio", 2.0)) {
			return fmt.Sprintf("Hot Planet: distance to star %.2f Ls (%.2f Star Rad)",
				math.Abs(body.SemiMajorAxis-parent.Radius)/LIGHT_SECOND, smaRadRatio)
		}
	}

	return ""
}

func (cs *CurrentSystemT) highInclination(body *ScanT) string {
	incl := math.Abs(body.OrbitalInclination)
	if incl >= 45.0 && incl <= 90.0+45.0 {
		return fmt.Sprintf("High Inclination: '%s' %+.1f&deg;", cs.BodyName(body.BodyName), body.OrbitalInclination)
	}
	return ""
}

func (cs *CurrentSystemT) fastSpinning(body *ScanT) string {
	if math.Abs(body.RotationPeriod) <= float64(sconf.Int32Def("criteria", "body rotation period", 1)) {
		return fmt.Sprintf("Fast Spinning: %.1f Hours", body.RotationPeriod/60/60)
	}

	return ""
}

func (cs *CurrentSystemT) highHeliumLevel(body *ScanT) string {
	if fnmatch.Match("*giant*", body.PlanetClass, fnmatch.FNM_IGNORECASE) {
		for _, atmo := range body.AtmosphereComposition {
			if atmo.Name == "Helium" && atmo.Percent >= float64(sconf.Float32Def("criteria", "min helium level", 29.0)) {
				return fmt.Sprintf("High Helium level: %.1f%%", atmo.Percent)
			}
		}
	}

	return ""
}
