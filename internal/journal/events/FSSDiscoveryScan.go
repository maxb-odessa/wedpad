package events

import (
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// FSSDiscoveryScan event structure
type FSSDiscoveryScanT struct {
	BodyCount     int       `mapstructure:"BodyCount"`
	NonBodyCount  int       `mapstructure:"NonBodyCount"`
	Progress      float64   `mapstructure:"Progress"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	SystemName    string    `mapstructure:"SystemName"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// FSSDiscoveryScan event handler
func (evh *EventHandler) FSSDiscoveryScan(eventData map[string]interface{}) {
	ev := new(FSSDiscoveryScanT)
	mapstructure.Decode(eventData, ev)

	cs := evh.CurrentSystem()

	cs.SetBodiesCount(ev.BodyCount)
	cs.SetNonBodiesCount(ev.NonBodyCount)

	type Bodies struct {
		Stars     int
		Bodies    int
		NonBodies int
	}

	starsCount := len(cs.Stars())

	B := &Bodies{Stars: starsCount, Bodies: ev.BodyCount - starsCount, NonBodies: ev.NonBodyCount}

	m := &msg.Message{
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
