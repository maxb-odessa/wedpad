package events

import (
	"fmt"
	"strings"

	"github.com/maxb-odessa/slog"

	"wedpad/internal/msg"
	"wedpad/internal/utils"
)

func (cs *CurrentSystemT) ShowStars() {

	baryCentres := cs.BaryCentres()
	currStars := cs.Stars()

	stars := make([]map[string]interface{}, 0)

	for _, s := range cs.StarsByNameSorted() { // array of *ScanT sorted by BodyName

		star := make(map[string]interface{})

		slog.Debug(9, "BODYNAME: '%s', CSNAME: '%s'", s.BodyName, cs.Name())
		tc := StarTypeColor(s.StarType)
		if tc.Color == "" {
			tc.Color = GuessColorByTemp(s.SurfaceTemperature)
		}

		if starName := utils.HTMLSafe(cs.BodyName(s.BodyName)); starName != "" {
			star["Name"] = starName
		} else {
			star["Name"] = "&#x2736;"
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
			star["Rings"] = fmt.Sprintf("%d / %.1f", rn, rr/LIGHT_SECOND)
		}

		var parents = []string{}
		for _, par := range s.Parents {
			if st, ok := currStars[par.Star]; ok {
				stName := utils.HTMLSafe(cs.BodyName(st.BodyName))
				if stName != "" {
					parents = append(parents, stName)
				} else {
					parents = append(parents, "&#x2736;")
				}
			}
			if _, ok := baryCentres[par.Null]; ok {
				parents = append(parents, "&#9741;")
			}
		}

		star["Parents"] = strings.Join(parents, "&nbsp;")

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

	type Bodies struct {
		Stars     int
		Bodies    int
		NonBodies int
	}

	starsCount := len(currStars)

	B := &Bodies{Stars: starsCount, Bodies: cs.BodiesCount() - starsCount, NonBodies: cs.NonBodiesCount()}

	m = &msg.Message{
		Target: msg.TARGET_STARS,
		Type:   msg.TYPE_BUTTON,
		Action: msg.ACTION_REPLACE,
		Data:   B,
	}

	m.Send()

	m = &msg.Message{
		Target: msg.TARGET_STARS,
		Type:   msg.TYPE_BUTTON,
		Action: msg.ACTION_ATTENTION,
		Data:   "",
	}

	m.Send()
}
