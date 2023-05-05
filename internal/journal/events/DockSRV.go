package events

import (
	"time"
)

// DockSRV event structure
type evDockSRV struct {
	ID               int       `mapstructure:"ID"`
	SrvType          string    `mapstructure:"SRVType"`
	SRVTypeLocalised string    `mapstructure:"SRVType_Localised"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// DockSRV event handler
func DockSRV(e interface{}) {
    // ev := e.(evDockSRV)
}

