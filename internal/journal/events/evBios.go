package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/danwakefield/fnmatch"
	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
	"github.com/mitchellh/mapstructure"
)

type BioT struct {
	Family         string     `mapstructure:"Family"`
	Name           string     `mapstructure:"Name"`
	Atmospheres    []string   `mapstructure:"Atmospheres"`
	BodiesRequired []string   `mapstructure:"BodiesRequired"`
	OnBodies       []string   `mapstructure:"OnBodies"`
	ColonyRangeM   int        `mapstructure:"ColonyRangeM"`
	GravityG       [2]float64 `mapstructure:"GravityG"`
	TemperatureK   [2]float64 `mapstructure:"TemperatureK"`
	StarsRequired  []string   `mapstructure:"StarsRequired"`
	Volcanism      []string   `mapstructure:"Volcanism"`
	ValueCr        int        `mapstructure:"ValueCr"`
	MinDistanceLs  int        `mapstructure:"MinDistanceLs"`
	Notes          string     `mapstructure:"Notes"`
}

type BiosT struct {
	bios []*BioT
}

func (b *BiosT) Init(bioData []byte) error {

	if len(bioData) == 0 {
		return errors.New("Empty BIO data set")
	}

	jsData := make([]interface{}, 0)
	if err := json.Unmarshal(bioData, &jsData); err != nil {
		return err
	}

	b.bios = make([]*BioT, 0)

	for _, bioData := range jsData {
		bio := new(BioT)
		if err := mapstructure.Decode(bioData, bio); err != nil {
			return fmt.Errorf("Decode BIO data failed: %s", err)
		}
		b.bios = append(b.bios, bio)
	}

	slog.Info("Loaded %d bio definitions", len(b.bios))

	return nil
}

// called once after FSSAllBodiesFound event
func (b *BiosT) Predict(cs *CurrentSystemT) map[string][2]string {

	predicted := make(map[string][2]string, 0)

	for _, planet := range cs.Planets() {

		// the planet has no bios on it
		if !bodyHasBios(cs, planet) {
			continue
		}

		// no bios guessed?
		bios := b.guessBios(cs, planet)
		if len(bios) == 0 {
			slog.Warn("Failed to guess BIO on '%s'", planet.BodyName)
			continue
		}

		// is acceptable? (DSS returned real bios, compare them with predicted families)
		if acceptBios, ok := cs.PlanetSignalsFound()[planet.BodyID]; ok {

			tempBios := make([]*BioT, 0)

			for _, accept := range acceptBios.Genuses {
				for _, bio := range bios {
					if accept.GenusLocalised == bio.Family {
						tempBios = append(tempBios, bio)
						//break
					}
				}
			}

			if len(tempBios) > 0 {
				bios = tempBios
			}

		}

		bioFams := make(map[string]bool, 0)
		bioHints := make([]string, 0)
		for _, bio := range bios {

			bioFams[bio.Family] = true

			if float32(bio.ValueCr/1_000_000.0) >= sconf.Float32Def("criteria", "min bio value", 0.0) {
				bioHints = append(bioHints, makeHint(bio))
			}

		}

		bioList := make([]string, 0)
		for f, _ := range bioFams {
			bioList = append(bioList, f)
		}

		sort.Strings(bioList)
		sort.Strings(bioHints)

		hintsTable := `<table class="hint-table">` +
			`<thead><tr>` +
			`<th>Variant</th>` +
			`<th>Pay(MCr)</th>` +
			`<th>Range(m)</th>` +
			`<th>Notes</th>` +
			`</tr></thead><tbody>` +
			strings.Join(bioHints, "") +
			`</tbody></table>`

		predicted[planet.BodyName] = [2]string{strings.Join(bioList, ", "), hintsTable}

	}

	return predicted
}

func makeHint(bio *BioT) string {
	return `<tr>` +
		`<td id="text-left">` + strings.ReplaceAll(bio.Name, " ", "&nbsp;") + `</td>` +
		`<td id="text-right">` + fmt.Sprintf("%.1f", float64(bio.ValueCr)/1_000_000.0) + `</td>` +
		`<td id="text-right">` + fmt.Sprintf("%d", bio.ColonyRangeM) + `</td>` +
		`<td id="text-left" style="width: 100%; font-size: smaller;">` + bio.Notes + `</td>` +
		`</tr>`
}

func (b *BiosT) getBio(name string) *BioT {

	for _, bio := range b.bios {
		if bio.Name == name {
			return bio
		}
	}

	return nil
}

func (b *BiosT) guessBios(cs *CurrentSystemT, ev *ScanT) []*BioT {

	pBios := make([]*BioT, 0)

	for _, bio := range b.bios {

		if !matchStrings(ev.Atmosphere, bio.Atmospheres) {
			continue
		}

		if !matchStrings(ev.PlanetClass, bio.OnBodies) {
			continue
		}

		if !inRange(ev.SurfaceTemperature, bio.TemperatureK) {
			continue
		}

		if !inRange(ev.SurfaceGravity/10.0, bio.GravityG) { // again, shoud divide by 10
			continue
		}

		if ev.DistanceFromArrivalLs < float64(bio.MinDistanceLs) {
			continue
		}

		if len(bio.StarsRequired) > 0 && !matchStars(cs, bio.StarsRequired) {
			continue
		}

		if len(bio.BodiesRequired) > 0 && !matchBodies(cs, bio.BodiesRequired) {
			continue
		}

		if len(bio.Volcanism) > 0 {
			if ev.Volcanism == "" || !matchStrings(ev.Volcanism, bio.Volcanism) {
				continue
			}
		}

		pBios = append(pBios, bio)

	}

	return pBios
}

func bodyHasBios(cs *CurrentSystemT, ev *ScanT) bool {
	if bs, ok := cs.PlanetSignalsCount()[ev.BodyID]; ok {
		for _, s := range bs.Signals {
			if s.Type == "$SAA_SignalType_Biological;" && s.Count > 0 {
				return true
			}
		}
	}

	return false
}

func matchStrings(what string, patts []string) bool {
	for _, patt := range patts {
		if fnmatch.Match(patt, what, fnmatch.FNM_IGNORECASE) {
			return true
		}
	}

	return false
}

func inRange(what float64, where [2]float64) bool {
	if what >= where[0] && what <= where[1] {
		return true
	}

	return false
}

func matchStars(cs *CurrentSystemT, starsRequired []string) bool {
	stars := cs.Stars()

	for _, star := range stars {
		if matchStrings(star.StarType, starsRequired) {
			return true
		}
	}

	return false
}

func matchBodies(cs *CurrentSystemT, bodiesRequired []string) bool {
	planets := cs.Planets()

	for _, planet := range planets {
		if matchStrings(planet.PlanetClass, bodiesRequired) {
			return true
		}
	}

	return false
}
