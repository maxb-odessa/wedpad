package events

// DockingCancelled event structure
type DockingCancelledT interface{}

// DockingCancelled event handler
func (evh *EventHandler) DockingCancelled(eventData map[string]interface{}) {
    // ev := new(DockingCancelledT)
    // mapstructure.Decode(eventData, ev)
}

