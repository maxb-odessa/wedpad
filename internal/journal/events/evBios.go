package events

import (
	"encoding/json"
	"errors"
	"fmt"

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
	TemperatureK   [2]int     `mapstructure:"TemperatureK"`
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

	bios := make([]string, 0)

	return bios
}
