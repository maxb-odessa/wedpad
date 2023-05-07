package events

// null event structure
type evnull interface{}

// null event handler
func (evHandler EventHandler) null(eventData map[string]interface{}) {
    // ev := new(evnull)
    // mapstructure.Decode(eventData, ev)
}

