package events

import (
	"time"
)

// FSSSignalDiscovered event structure
type FSSSignalDiscoveredT struct {
	IsStation                bool      `mapstructure:"IsStation"`
	SignalName               string    `mapstructure:"SignalName"`
	SignalNameLocalised      string    `mapstructure:"SignalName_Localised"`
	SpawningFaction          string    `mapstructure:"SpawningFaction"`
	SpawningFactionLocalised string    `mapstructure:"SpawningFaction_Localised"`
	SpawningState            string    `mapstructure:"SpawningState"`
	SpawningStateLocalised   string    `mapstructure:"SpawningState_Localised"`
	SystemAddress            int       `mapstructure:"SystemAddress"`
	ThreatLevel              int       `mapstructure:"ThreatLevel"`
	TimeRemaining            float64   `mapstructure:"TimeRemaining"`
	UssType                  string    `mapstructure:"USSType"`
	USSTypeLocalised         string    `mapstructure:"USSType_Localised"`
	Event                    string    `mapstructure:"event"`
	Timestamp                time.Time `mapstructure:"timestamp"`
}

// FSSSignalDiscovered event handler
func (evh *EventHandler) FSSSignalDiscovered(eventData map[string]interface{}) {
    // ev := new(FSSSignalDiscoveredT)
    // mapstructure.Decode(eventData, ev)
}

