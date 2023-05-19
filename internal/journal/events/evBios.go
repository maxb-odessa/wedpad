package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/danwakefield/fnmatch"
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

		bioFams := make(map[string]bool, 0)
		bioHints := make([]string, 0)
		for _, bio := range bios {

			bioFams[bio.Family] = true

			hint := fmt.Sprintf(`<tr><td>%s</td><td style="text-align: right;">%.1f MCr</td>`, bio.Name, float64(bio.ValueCr)/1_000_000.0)
			if bio.Notes != "" {
				hint += "<td>" + bio.Notes + "</td>"
			}
			hint += "</tr>"

			bioHints = append(bioHints, hint)

		}

		bioList := make([]string, 0)
		for f, _ := range bioFams {
			bioList = append(bioList, f)
		}

		sort.Strings(bioList)
		sort.Strings(bioHints)

		predicted[planet.BodyName] = [2]string{strings.Join(bioList, ", "), strings.Join(bioHints, "")}

	}

	return predicted
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
		/*
			if !matchStars(cs.Stars(), bio.StarsRequired) {
				continue
			}

			if !matchBodies(cs.Planets(), bio.BodiesRequired) {
				continue
			}
		*/
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
