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
func (evHandler EventHandler) FSSDiscoveryScan(eventData map[string]interface{}) {
	ev := new(FSSDiscoveryScanT)
	mapstructure.Decode(eventData, ev)

	type Bodies struct {
		Bodies    int
		NonBodies int
	}

	B := &Bodies{ev.BodyCount, ev.NonBodyCount}

	md := &msg.Message{
		Target: msg.TARGET_SYSTEM,
		Type:   msg.TYPE_BUTTON,
		Action: msg.ACTION_REPLACE,
		Data:   B,
	}

	md.Send()

	ma := &msg.Message{
		Target: msg.TARGET_SYSTEM,
		Type:   msg.TYPE_BUTTON,
		Action: msg.ACTION_ATTENTION,
		Data:   "",
	}

	ma.Send()
}
