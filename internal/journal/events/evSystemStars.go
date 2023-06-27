package events

import (
	"fmt"
	"sort"

	"github.com/fvbommel/sortorder"
	"github.com/maxb-odessa/slog"

	"wedpad/internal/msg"
	"wedpad/internal/utils"
)

func (cs *CurrentSystemT) ShowStars() {

	stars := make([]map[string]interface{}, 0)

	currSysStars := cs.StarsByName()

	var keys []string
	for k, _ := range currSysStars {
		keys = append(keys, k)
	}

	sort.Strings(sortorder.Natural(keys))

	for _, starName := range keys {

		star := make(map[string]interface{})

		if starName == cs.MainStarName() {
			star["MainStar"] = true
		} else {
			star["MainStar"] = false
		}

		s := currSysStars[starName]

		slog.Debug(9, "BODYNAME: '%s', CSNAME: '%s'", s.BodyName, cs.Name())
		star["Barycenter"] = false
		star["Name"] = utils.HTMLSafe(cs.BodyName(s.BodyName))
		tc := StarTypeColor(s.StarType)
		if tc.Color == "" {
			tc.Color = GuessColorByTemp(s.SurfaceTemperature)
		}
		star["Type"] = tc
		star["Subclass"] = s.Subclass
		star["Luminosity"] = s.Luminosity
		if s.DistanceFromArrivalLs > 0.0 {
			star["DistanceLs"] = Num(s.DistanceFromArrivalLs)
		}
		star["RadiusS"] = Num(s.Radius / SOLAR_RADIUS)
		star["MassS"] = Num(s.StellarMass)
		star["TemperatureK"] = Num(s.SurfaceTemperature)
		star["TemperatureColor"] = GuessColorByTemp(s.SurfaceTemperature)
		star["OrbitalPeriodD"] = Num(s.OrbitalPeriod / SECONDS_IN_DAY)
		star["Eccentricity"] = Num(s.Eccentricity)
		star["Discovered"] = s.WasDiscovered
		rn, rr := CalcRings(s)
		if rn > 0 {
			star["Rings"] = fmt.Sprintf("%d/%.1f", rn, rr/LIGHT_SECOND)
		}

		stars = append(stars, star)
	}

	// send data to system view
	m := &msg.Message{
		Action: msg.ACTION_REPLACE,
		Target: msg.TARGET_STARS,
		Type:   msg.TYPE_VIEW,
		Data:   stars,
	}

	m.Send()

	// switch to system view
	m = &msg.Message{
		Action: msg.ACTION_ATTENTION,
		Target: msg.TARGET_STARS,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}

	m.Send()
}
