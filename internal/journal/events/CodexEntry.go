package events

import (
	"fmt"
	"strings"
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// CodexEntry event structure
type CodexEntryT struct {
	BodyID                      int       `mapstructure:"BodyID"`
	Category                    string    `mapstructure:"Category"`
	CategoryLocalised           string    `mapstructure:"Category_Localised"`
	EntryID                     int       `mapstructure:"EntryID"`
	IsNewEntry                  bool      `mapstructure:"IsNewEntry"`
	Latitude                    float64   `mapstructure:"Latitude"`
	Longitude                   float64   `mapstructure:"Longitude"`
	Name                        string    `mapstructure:"Name"`
	NameLocalised               string    `mapstructure:"Name_Localised"`
	NearestDestination          string    `mapstructure:"NearestDestination"`
	NearestDestinationLocalised string    `mapstructure:"NearestDestination_Localised"`
	Region                      string    `mapstructure:"Region"`
	RegionLocalised             string    `mapstructure:"Region_Localised"`
	SubCategory                 string    `mapstructure:"SubCategory"`
	SubCategoryLocalised        string    `mapstructure:"SubCategory_Localised"`
	System                      string    `mapstructure:"System"`
	SystemAddress               int       `mapstructure:"SystemAddress"`
	VoucherAmount               int       `mapstructure:"VoucherAmount"`
	Event                       string    `mapstructure:"event"`
	Timestamp                   time.Time `mapstructure:"timestamp"`
}

// CodexEntry event handler
func (evh *EventHandler) CodexEntry(eventData map[string]interface{}) {
	ev := new(CodexEntryT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()

	isNew := ""
	if !ev.IsNewEntry {
		isNew = " (new)"
	}

	valueCr := 0
	bioName := strings.Split(ev.NameLocalised, " - ")
	if b := cs.bios.getBio(bioName[0]); b != nil {
		valueCr = b.ValueCr
	}

	text := fmt.Sprintf("Codex%s: %s: %s, region: %s, value: %.2fMCr",
		isNew,
		ev.SubCategoryLocalised,
		ev.NameLocalised,
		ev.RegionLocalised,
		float32(valueCr/1_000_000.0),
	)

	m := &msg.Message{
		Type:   msg.TYPE_VIEW,
		Target: msg.TARGET_LOG,
		Action: msg.ACTION_APPEND,
		Data:   text,
	}
	m.Send()

	m = &msg.Message{
		Type:   msg.TYPE_BUTTON,
		Target: msg.TARGET_LOG,
		Action: msg.ACTION_ATTENTION,
		Data:   "",
	}
	m.Send()

}
