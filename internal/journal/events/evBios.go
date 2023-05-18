package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/danwakefield/fnmatch"
	"github.com/maxb-odessa/slog"
	"github.com/mitchellh/mapstructure"
)

type BioT struct {
	Family         string     `mapstructure:"Family"`
	Name           string     `mapstructure:"Type"`
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
			fmt.Errorf("Decode BIO data failed: %s", err)
			return err
		}
		b.bios = append(b.bios, bio)
	}

	slog.Info("Loaded %d bio definitions", len(b.bios))

	return nil
}

func (b *BiosT) predictedBios(ev *ScanT) []string {

	pBios := make(map[string]bool, 0)

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
			if !matchStars(evStars(), bio.StarsRequired) {
				continue
			}

			if !matchBodies(ev.StarType, bio.StarsRequired) {
				continue
			}
		*/
		pBios[bio.Family] = true
	}

	bioList := make([]string, 0)
	for k, _ := range pBios {
		bioList = append(bioList, k)
	}

	sort.Strings(bioList)
	return bioList
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
