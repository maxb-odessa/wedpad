package events

import (
	"time"
)

// LaunchSRV event structure
type evLaunchSRV struct {
	ID               int       `mapstructure:"ID"`
	Loadout          string    `mapstructure:"Loadout"`
	PlayerControlled bool      `mapstructure:"PlayerControlled"`
	SrvType          string    `mapstructure:"SRVType"`
	SRVTypeLocalised string    `mapstructure:"SRVType_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// LaunchSRV event handler
func (evHandler EventHandler) LaunchSRV(eventData map[string]interface{}) {
    // ev := new(evLaunchSRV)
    // mapstructure.Decode(eventData, ev)
}

