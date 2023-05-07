package events

// Passengers event structure
type PassengersT interface{}

// Passengers event handler
func (evHandler EventHandler) Passengers(eventData map[string]interface{}) {
    // ev := new(PassengersT)
    // mapstructure.Decode(eventData, ev)
}

