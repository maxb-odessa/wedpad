package events

import (
	"sort"
	"wedpad/internal/msg"
)

func (cs *CurrentSystemT) ShowStars() {

	stars := make([]map[string]interface{}, 0)

	var keys []int
	for k, _ := range cs.Stars() {
		keys = append(keys, k)
	}

	for k, _ := range cs.BaryCentres() {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	currSysStars := cs.Stars()

	for _, id := range keys {

		star := make(map[string]interface{})

		if id == cs.MainStarID() {
			star["MainStar"] = true
		} else {
			star["MainStar"] = false
		}

		// a Star, not a BaryCentre
		if s, ok := currSysStars[id]; ok {
			star["Barycenter"] = false
			star["Name"] = BodyName(s.BodyName, cs.Name())
			star["Type"] = StarTypeColor(s.StarType)
			star["Subclass"] = s.Subclass
			star["Luminosity"] = s.Luminosity
			star["DistanceLs"] = Num(s.DistanceFromArrivalLs)
			star["RadiusS"] = Num(s.Radius / SOLAR_RADIUS)
			star["MassS"] = Num(s.StellarMass)
			star["TemperatureK"] = Num(s.SurfaceTemperature)
			star["OrbitalPeriodD"] = Num(s.OrbitalPeriod / SECONDS_IN_DAY)
			star["Discovered"] = s.WasDiscovered
			rn, rr := CalcRings(s)
			if rn > 0 {
				star["RingsNum"] = rn
				star["RingsRadiusLs"] = Num(rr / LIGHT_SECOND)
			}
		} else {
			star["Barycenter"] = true
			// TODO get it from star parents star["Name"] = BodyName(s.BodyName, cs.Name())
			star["Name"] = BodyName("+", cs.Name())
		}

		stars = append(stars, star)
	}

	// send data to system view
	mv := &msg.Message{
		Action: msg.ACTION_REPLACE,
		Target: msg.TARGET_SYSTEM,
		Type:   msg.TYPE_VIEW,
		Data: map[string]interface{}{
			"SystemName": cs.Name(),
			"Data":       stars,
		},
	}

	mv.Send()

	// switch to system view
	mb := &msg.Message{
		Action: msg.ACTION_ATTENTION,
		Target: msg.TARGET_SYSTEM,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}

	mb.Send()
}
