package events

// BuyDrones event structure
type BuyDronesT interface{}

// BuyDrones event handler
func (evHandler EventHandler) BuyDrones(eventData map[string]interface{}) {
    // ev := new(BuyDronesT)
    // mapstructure.Decode(eventData, ev)
}

