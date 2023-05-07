package events

// CargoDepot event structure
type CargoDepotT interface{}

// CargoDepot event handler
func (evHandler EventHandler) CargoDepot(eventData map[string]interface{}) {
    // ev := new(CargoDepotT)
    // mapstructure.Decode(eventData, ev)
}

