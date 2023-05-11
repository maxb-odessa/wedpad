package events

// SystemsShutdown event structure
type SystemsShutdownT interface{}

// SystemsShutdown event handler
func (evh *EventHandler) SystemsShutdown(eventData map[string]interface{}) {
    // ev := new(SystemsShutdownT)
    // mapstructure.Decode(eventData, ev)
}

