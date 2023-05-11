package events

// SellDrones event structure
type SellDronesT interface{}

// SellDrones event handler
func (evh *EventHandler) SellDrones(eventData map[string]interface{}) {
    // ev := new(SellDronesT)
    // mapstructure.Decode(eventData, ev)
}

