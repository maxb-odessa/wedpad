package events

// BuyExplorationData event structure
type BuyExplorationDataT interface{}

// BuyExplorationData event handler
func (evHandler EventHandler) BuyExplorationData(eventData map[string]interface{}) {
    // ev := new(BuyExplorationDataT)
    // mapstructure.Decode(eventData, ev)
}

