package events

import (
	"time"
	"wedpad/internal/msg"
)

// NavRouteClear event structure
type NavRouteClearT struct {
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// NavRouteClear event handler
func (evh *EventHandler) NavRouteClear(eventData map[string]interface{}) {
	// ev := new(NavRouteClearT)
	// mapstructure.Decode(eventData, ev)

	m := &msg.Message{
		Action: msg.ACTION_CLEAN,
		Target: msg.TARGET_TOP,
		Type:   msg.TYPE_VIEW,
		Data:   "",
	}

	m.Send()
}
