package events

// Died event structure
type DiedT interface{}

// Died event handler
func (evHandler EventHandler) Died(eventData map[string]interface{}) {
    // ev := new(DiedT)
    // mapstructure.Decode(eventData, ev)
}

