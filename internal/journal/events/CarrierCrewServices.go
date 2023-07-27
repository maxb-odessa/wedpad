package events

import (
	"time"
)

// CarrierCrewServices event structure
type CarrierCrewServicesT struct {
	CarrierID int       `mapstructure:"CarrierID"`
	CrewName  string    `mapstructure:"CrewName"`
	CrewRole  string    `mapstructure:"CrewRole"`
	Operation string    `mapstructure:"Operation"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CarrierCrewServices event handler
func (evh *EventHandler) CarrierCrewServices(eventData map[string]interface{}) {
    // ev := new(CarrierCrewServicesT)
    // mapstructure.Decode(eventData, ev)
    // cs := evh.CurrentSystem()
}

