package events

import (
	"time"
)

// CarrierFinance event structure
type CarrierFinanceT struct {
	AvailableBalance int       `mapstructure:"AvailableBalance"`
	CarrierBalance   int       `mapstructure:"CarrierBalance"`
	CarrierID        int       `mapstructure:"CarrierID"`
	ReserveBalance   int       `mapstructure:"ReserveBalance"`
	ReservePercent   int       `mapstructure:"ReservePercent"`
	TaxRateRearm     int       `mapstructure:"TaxRate_rearm"`
	TaxRateRefuel    int       `mapstructure:"TaxRate_refuel"`
	TaxRateRepair    int       `mapstructure:"TaxRate_repair"`
	Event            string    `mapstructure:"event"`
	Timestamp        time.Time `mapstructure:"timestamp"`
}

// CarrierFinance event handler
func (evh *EventHandler) CarrierFinance(eventData map[string]interface{}) {
    // ev := new(CarrierFinanceT)
    // mapstructure.Decode(eventData, ev)
}

