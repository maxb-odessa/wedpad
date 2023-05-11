package events

import (
	"time"
)

// LaunchSRV event structure
type LaunchSRVT struct {
	ID               int       `mapstructure:"ID"`
	Loadout          string    `mapstructure:"Loadout"`
	PlayerControlled bool      `mapstructure:"PlayerControlled"`
	SrvType          string    `mapstructure:"SRVType"`
	SRVTypeLocalised string    `mapstructure:"SRVType_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// LaunchSRV event handler
func (evh *EventHandler) LaunchSRV(eventData map[string]interface{}) {
    // ev := new(LaunchSRVT)
    // mapstructure.Decode(eventData, ev)
}

