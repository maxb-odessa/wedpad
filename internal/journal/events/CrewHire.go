package events

// CrewHire event structure
type CrewHireT interface{}

// CrewHire event handler
func (evHandler EventHandler) CrewHire(eventData map[string]interface{}) {
    // ev := new(CrewHireT)
    // mapstructure.Decode(eventData, ev)
}

