package events

import (
	"fmt"
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// MaterialCollected event structure
type MaterialCollectedT struct {
	Category      string    `mapstructure:"Category"`
	Count         int       `mapstructure:"Count"`
	Name          string    `mapstructure:"Name"`
	NameLocalised string    `mapstructure:"Name_Localised"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// MaterialCollected event handler
func (evh *EventHandler) MaterialCollected(eventData map[string]interface{}) {
	ev := new(MaterialCollectedT)
	mapstructure.Decode(eventData, ev)

	m := &msg.Message{
		Action: msg.ACTION_APPEND,
		Target: msg.TARGET_LOG,
		Type:   msg.TYPE_VIEW,
		Data:   fmt.Sprintf("Collected <b>%d</b> units of %s <b>%s</b>", ev.Count, ev.Category, ev.Name),
	}

	m.Send()
}
