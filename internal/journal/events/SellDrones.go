package events

// SellDrones event structure
type SellDronesT interface{}

// SellDrones event handler
func (evHandler EventHandler) SellDrones(eventData map[string]interface{}) {
    // ev := new(SellDronesT)
    // mapstructure.Decode(eventData, ev)
}

