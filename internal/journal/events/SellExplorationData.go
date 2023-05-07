package events

// SellExplorationData event structure
type SellExplorationDataT interface{}

// SellExplorationData event handler
func (evHandler EventHandler) SellExplorationData(eventData map[string]interface{}) {
    // ev := new(SellExplorationDataT)
    // mapstructure.Decode(eventData, ev)
}

