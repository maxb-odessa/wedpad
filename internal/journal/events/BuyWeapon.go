package events

import (
	"time"
)

// BuyWeapon event structure
type evBuyWeapon struct {
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
func BuyWeapon(e interface{}) {
    // ev := e.(evBuyWeapon)
}

