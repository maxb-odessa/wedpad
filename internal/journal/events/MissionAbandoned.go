package events

// MissionAbandoned event structure
type MissionAbandonedT interface{}

// MissionAbandoned event handler
func (evHandler EventHandler) MissionAbandoned(eventData map[string]interface{}) {
    // ev := new(MissionAbandonedT)
    // mapstructure.Decode(eventData, ev)
}

