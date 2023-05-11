package events

// PayFines event structure
type PayFinesT interface{}

// PayFines event handler
func (evh *EventHandler) PayFines(eventData map[string]interface{}) {
    // ev := new(PayFinesT)
    // mapstructure.Decode(eventData, ev)
}

