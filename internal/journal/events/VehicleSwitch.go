package events

// VehicleSwitch event structure
type VehicleSwitchT interface{}

// VehicleSwitch event handler
func (evHandler EventHandler) VehicleSwitch(eventData map[string]interface{}) {
    // ev := new(VehicleSwitchT)
    // mapstructure.Decode(eventData, ev)
}

