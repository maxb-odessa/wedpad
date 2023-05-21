package events

import (
	"fmt"
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// ScanOrganic event structure
type ScanOrganicT struct {
	Body             int       `mapstructure:"Body"`
	Genus            string    `mapstructure:"Genus"`
	GenusLocalised   string    `mapstructure:"Genus_Localised"`
	ScanType         string    `mapstructure:"ScanType"`
	Species          string    `mapstructure:"Species"`
	SpeciesLocalised string    `mapstructure:"Species_Localised"`
	SystemAddress    int       `mapstructure:"SystemAddress"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// ScanOrganic event handler
func (evh *EventHandler) ScanOrganic(eventData map[string]interface{}) {
	ev := new(ScanOrganicT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()

	bio := cs.bios.getBio(ev.SpeciesLocalised)

	m := &msg.Message{
		Type:   msg.TYPE_VIEW,
		Target: msg.TARGET_LOG,
		Action: msg.ACTION_APPEND,
		Data:   fmt.Sprintf("Biological "+ev.ScanType+": "+ev.SpeciesLocalised+", diversity: %d meters", bio.ColonyRangeM),
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
