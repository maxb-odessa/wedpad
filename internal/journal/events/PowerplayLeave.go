package events

// PowerplayLeave event structure
type PowerplayLeaveT interface{}

// PowerplayLeave event handler
func (evHandler EventHandler) PowerplayLeave(eventData map[string]interface{}) {
    // ev := new(PowerplayLeaveT)
    // mapstructure.Decode(eventData, ev)
}

