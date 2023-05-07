package events

import (
	"time"
)

// SetUserShipName event structure
type evSetUserShipName struct {
	Ship         string    `mapstructure:"Ship"`
	ShipID       int       `mapstructure:"ShipID"`
	UserShipID   string    `mapstructure:"UserShipId"`
	UserShipName string    `mapstructure:"UserShipName"`
	Event        string    `mapstructure:"event"`
	Timestamp    time.Time `mapstructure:"timestamp"`
}

// SetUserShipName event handler
func (evHandler EventHandler) SetUserShipName(eventData map[string]interface{}) {
    // ev := new(evSetUserShipName)
    // mapstructure.Decode(eventData, ev)
}
