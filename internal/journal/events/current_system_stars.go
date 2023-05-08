package events

import (
	"sort"
	"wedpad/internal/msg"
)

func (cs *CurrentSystemT) ShowStars() {

	stars := make([]map[string]interface{}, 0)

	var keys []int
	for k, _ := range CurrentSystem.Stars {
		keys = append(keys, k)
	}

	for k, _ := range CurrentSystem.BaryCentres {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, id := range keys {

		star := make(map[string]interface{})

		// a Star, not a BaryCentre
		if s, ok := CurrentSystem.Stars[id]; ok {
			star["Barycenter"] = false
			star["Links"] = `*`
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
				star["RingsRadiusLs"] = Num(rr)
			}
		} else {
			star["Barycenter"] = true
			star["Links"] = `--[`
		}

		stars = append(stars, star)
	}

	type withSystemName struct {
		SystemName string
		Data       interface{}
	}
	Data := &withSystemName{
		SystemName: CurrentSystem.String(),
		Data:       stars,
	}

	mv := &msg.Message{
		Action: msg.ACTION_REPLACE,
		Target: msg.TARGET_SYSTEM,
		Type:   msg.TYPE_VIEW,
		Data:   Data,
	}

	mv.Send()

	mb := &msg.Message{
		Action: msg.ACTION_ATTENTION,
		Target: msg.TARGET_SYSTEM,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}

	mb.Send()
}
