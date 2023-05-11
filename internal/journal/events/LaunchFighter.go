package events

// LaunchFighter event structure
type LaunchFighterT interface{}

// LaunchFighter event handler
func (evh *EventHandler) LaunchFighter(eventData map[string]interface{}) {
    // ev := new(LaunchFighterT)
    // mapstructure.Decode(eventData, ev)
}

