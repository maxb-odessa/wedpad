package events

import (
	"time"
)

// CarrierJumpRequest event structure
type CarrierJumpRequestT struct {
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
func (evh *EventHandler) CarrierJumpRequest(eventData map[string]interface{}) {
    // ev := new(CarrierJumpRequestT)
    // mapstructure.Decode(eventData, ev)
}

