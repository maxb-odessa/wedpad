package events

// SellSuit event structure
type SellSuitT interface{}

// SellSuit event handler
func (evh *EventHandler) SellSuit(eventData map[string]interface{}) {
    // ev := new(SellSuitT)
    // mapstructure.Decode(eventData, ev)
}

