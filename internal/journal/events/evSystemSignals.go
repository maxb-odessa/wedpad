package events

import (
	"wedpad/internal/msg"
)

type SignalT struct {
	Name        string
	Type        string
	Description string
	Hint        string
}

func (cs *CurrentSystemT) ShowSignals() {

	signals := make([]*SignalT, 0)

	// add system signals (phenomena, stations, etc)
	for _, sig := range cs.Signals() {

		s := new(SignalT)

		if sig.IsStation {
			s.Name = "&nabla;"
			s.Type = `<font color="cyan">Station</font>`
		} else {
			// TODO for now: collect all sig.SignalName variants
			s.Name = "&sect;"
			s.Type = `<font color="yellow">Phenomena</font>`
		}

		if sig.SignalNameLocalised != "" {
			s.Description = sig.SignalNameLocalised + "<br>(" + sig.SignalName + ")"
		} else {
			s.Description = sig.SignalName
		}

		signals = append(signals, s)
	}

	// predict and add bodies BIO signals
	for name, sigs := range cs.bios.Predict(cs) {
		s := &SignalT{
			Name:        cs.BodyName(name),
			Type:        `<font color="limegreen">Biological</font>`,
			Description: sigs[0],
			Hint:        sigs[1],
		}
		signals = append(signals, s)
	}

	sigsCnt := len(signals)
	if sigsCnt > 0 {

		m := &msg.Message{
			Target: msg.TARGET_SIGNALS,
			Type:   msg.TYPE_VIEW,
			Action: msg.ACTION_REPLACE,
			Data:   signals,
		}
		m.Send()

		m = &msg.Message{
			Target: msg.TARGET_SIGNALS,
			Type:   msg.TYPE_BUTTON,
			Action: msg.ACTION_REPLACE,
			Data:   sigsCnt,
		}
		m.Send()

		m = &msg.Message{
			Target: msg.TARGET_SIGNALS,
			Type:   msg.TYPE_BUTTON,
			Action: msg.ACTION_ATTENTION,
			Data:   "",
		}
		m.Send()
	}
}
