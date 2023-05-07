package events

// HeatDamage event structure
type HeatDamageT interface{}

// HeatDamage event handler
func (evHandler EventHandler) HeatDamage(eventData map[string]interface{}) {
    // ev := new(HeatDamageT)
    // mapstructure.Decode(eventData, ev)
}

