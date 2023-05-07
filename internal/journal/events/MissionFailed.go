package events

// MissionFailed event structure
type MissionFailedT interface{}

// MissionFailed event handler
func (evHandler EventHandler) MissionFailed(eventData map[string]interface{}) {
    // ev := new(MissionFailedT)
    // mapstructure.Decode(eventData, ev)
}

