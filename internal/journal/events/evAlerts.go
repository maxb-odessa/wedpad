package events

import (
	"fmt"
	"wedpad/internal/msg"
)

type Alert struct {
	Text  string
	Level string
}

var alerts map[string]*Alert

const (
	ALERT_FUEL = "fuel"

/*
	not implemented yet

ALERT_HEAT      = "heat"
ALERT_DAMAGE    = "damage"
ALERT_PROXIMITY = "proximity"
ALERT_IMPACT    = "impact"
*/
)

const (
	ALERT_LEVEL_CRIT = "crit"
	ALERT_LEVEL_WARN = "warn"
	ALERT_LEVEL_INFO = "info"
)

func alertAdd(id string, level string, text string) {
	alerts[id] = &Alert{
		Text:  text,
		Level: level,
	}
}

func alertDel(id string) bool {
	if _, ok := alerts[id]; ok {
		delete(alerts, id)
		return true
	}
	return false
}

func alertShow() {

	m := &msg.Message{
		Target: msg.TARGET_BOTTOM,
		Action: msg.ACTION_REPLACE,
		Type:   msg.TYPE_VIEW,
		Data:   alerts,
	}

	m.Send()
}

func AlertFuel(fl float64) {
	alertDel(ALERT_FUEL)
	if fl < 10.0 {
		level := ALERT_LEVEL_INFO
		if fl < 5.0 {
			level = ALERT_LEVEL_WARN
			if fl < 2.0 {
				level = ALERT_LEVEL_CRIT
			}
		}
		alertAdd(ALERT_FUEL, level, fmt.Sprintf("Fuel level is %.1f tons", fl))
	}
	alertShow()
}
