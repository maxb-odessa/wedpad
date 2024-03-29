package events

import (
	"time"
)

// BuyWeapon event structure
type BuyWeaponT struct {
	Class         int           `mapstructure:"Class"`
	Name          string        `mapstructure:"Name"`
	NameLocalised string        `mapstructure:"Name_Localised"`
	Price         int           `mapstructure:"Price"`
	SuitModuleID  int           `mapstructure:"SuitModuleID"`
	WeaponMods    []interface{} `mapstructure:"WeaponMods"`
	Event         string        `mapstructure:"event"`
	Timestamp     time.Time     `mapstructure:"timestamp"`
}

// BuyWeapon event handler
func (evh *EventHandler) BuyWeapon(eventData map[string]interface{}) {
    // ev := new(BuyWeaponT)
    // mapstructure.Decode(eventData, ev)
}

