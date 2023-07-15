package events

import (
	"fmt"
	"time"
	"wedpad/internal/msg"

	"github.com/mitchellh/mapstructure"
)

// FuelScoop event structure
type FuelScoopT struct {
	Scooped   float64   `mapstructure:"Scooped"`
	Total     float64   `mapstructure:"Total"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// FuelScoop event handler
func (evh *EventHandler) FuelScoop(eventData map[string]interface{}) {

	ev := new(FuelScoopT)
	mapstructure.Decode(eventData, ev)

	m := &msg.Message{
		Target: msg.TARGET_LOG,
		Action: msg.ACTION_APPEND,
		Type:   msg.TYPE_VIEW,
		Data:   fmt.Sprintf("Fuel Scooped: %.1ft, Level: %.1ft", ev.Scooped, ev.Total),
	}
	m.Send()

}
