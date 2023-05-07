package events

// DockingTimeout event structure
type DockingTimeoutT interface{}

// DockingTimeout event handler
func (evHandler EventHandler) DockingTimeout(eventData map[string]interface{}) {
    // ev := new(DockingTimeoutT)
    // mapstructure.Decode(eventData, ev)
}

