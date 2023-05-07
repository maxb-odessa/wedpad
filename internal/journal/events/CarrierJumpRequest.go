package events

import (
	"time"
)

// CarrierJumpRequest event structure
type evCarrierJumpRequest struct {
	Body          string    `mapstructure:"Body"`
	BodyID        int       `mapstructure:"BodyID"`
	CarrierID     int       `mapstructure:"CarrierID"`
	DepartureTime time.Time `mapstructure:"DepartureTime"`
	SystemAddress int       `mapstructure:"SystemAddress"`
	SystemName    string    `mapstructure:"SystemName"`
	Event         string    `mapstructure:"event"`
	Timestamp     time.Time `mapstructure:"timestamp"`
}

// CarrierJumpRequest event handler
func (evHandler EventHandler) CarrierJumpRequest(eventData map[string]interface{}) {
    // ev := new(evCarrierJumpRequest)
    // mapstructure.Decode(eventData, ev)
}
