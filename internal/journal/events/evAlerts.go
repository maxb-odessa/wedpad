package events

import (
	"fmt"
	"wedpad/internal/msg"
	"wedpad/internal/sound"
)

type AlertMsgT struct {
	Text  string
	Level string
}

type AlertT struct {
	alerts map[string]*AlertMsgT
	sound  *sound.SoundT
}

func (al *AlertT) Init(snd *sound.SoundT) error {
	al.alerts = make(map[string]*AlertMsgT)
	al.sound = snd
	return nil
}

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

func (al *AlertT) alertAdd(id string, level string, text string) {
	al.alerts[id] = &AlertMsgT{
		Text:  text,
		Level: level,
	}
}

func (al *AlertT) alertDel(id string) bool {
	if _, ok := al.alerts[id]; ok {
		delete(al.alerts, id)
		return true
	}
	return false
}

func (al *AlertT) alertShow() {

	m := &msg.Message{
		Target: msg.TARGET_BOTTOM,
		Action: msg.ACTION_REPLACE,
		Type:   msg.TYPE_VIEW,
		Data:   al.alerts,
	}

	m.Send()
}

func (al *AlertT) AlertFuel(fl float64) {

	al.alertDel(ALERT_FUEL)

	if fl < 10.0 {

		level := ALERT_LEVEL_INFO
		al.sound.Play("drop")

		if fl < 5.0 {

			level = ALERT_LEVEL_WARN
			al.sound.Play("drop")

			if fl < 2.0 {
				level = ALERT_LEVEL_CRIT
				al.sound.Play("drop")
			}

		}

		al.alertAdd(ALERT_FUEL, level, fmt.Sprintf("Fuel level is %.1f tons", fl))
	}

	al.alertShow()
}
