package events

import (
	"wedpad/internal/msg"
)

type AlertMsgT struct {
	Text  string
	Level string
}

type AlertT struct {
	alerts map[string]*AlertMsgT
}

func (al *AlertT) Init() error {
	al.alerts = make(map[string]*AlertMsgT)
	return nil
}

const (
	ALERT_LEVEL_NONE = "none"
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

func (al *AlertT) Alert(id string, level string, text string) {

	if level == ALERT_LEVEL_NONE {
		al.alertDel(id)
	} else {
		al.alertAdd(id, level, text)
	}

	al.alertShow()
}
