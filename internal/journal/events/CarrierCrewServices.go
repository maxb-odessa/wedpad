package events

import (
	"time"
)

// CarrierCrewServices event structure
type evCarrierCrewServices struct {
	CarrierID int       `mapstructure:"CarrierID"`
	CrewName  string    `mapstructure:"CrewName"`
	CrewRole  string    `mapstructure:"CrewRole"`
	Operation string    `mapstructure:"Operation"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// CarrierCrewServices event handler
func CarrierCrewServices(e interface{}) {
    // ev := e.(evCarrierCrewServices)
}

