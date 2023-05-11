package events

// NewCommander event structure
type NewCommanderT interface{}

// NewCommander event handler
func (evh *EventHandler) NewCommander(eventData map[string]interface{}) {
    // ev := new(NewCommanderT)
    // mapstructure.Decode(eventData, ev)
}

