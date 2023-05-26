package events

import (
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

	isNew := ""
	if !ev.IsNewEntry {
		isNew = " (new)"
	}

	text := "Codex" + isNew + ": " + ev.SubCategoryLocalised + ": <b>" + ev.NameLocalised + "</b>, region: " + ev.RegionLocalised

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
