package events

// BuyTradeData event structure
type BuyTradeDataT interface{}

// BuyTradeData event handler
func (evh *EventHandler) BuyTradeData(eventData map[string]interface{}) {
    // ev := new(BuyTradeDataT)
    // mapstructure.Decode(eventData, ev)
}

