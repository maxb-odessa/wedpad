package events

import (
	"wedpad/internal/msg"
	"wedpad/internal/utils"

	"github.com/danwakefield/fnmatch"
)

type SignalT struct {
	Name        string
	Type        string
	Description string
	Hints       string
}

func (cs *CurrentSystemT) ShowSignals() {

	signals := make([]*SignalT, 0)

	// add system signals (phenomena, stations, etc)
	for _, sig := range cs.Signals() {

		s := new(SignalT)

		sigName := utils.HTMLSafe(sig.SignalName)
		sigNameLoc := utils.HTMLSafe(sig.SignalNameLocalised)

		// TODO for now: collect all sig.SignalName variants
		if sig.IsStation {
			// a carrier should have its sign at the end, eg: FANCY NAME X7Y-ZD9
			if fnmatch.Match("* [A-Z0-9][A-Z0-9][A-Z0-9]-[A-Z0-9][A-Z0-9][A-Z0-9]", sig.SignalName, 0) {
				s.Name = "&nabla;"
				s.Type = `<font color="#FF6600">Carrier</font>`
			} else {
				s.Name = "&#8859;"
				s.Type = `<font color="cyan">Station</font>`
			}
		} else if fnmatch.Match("$Fixed_Event_Life*", sigName, 0) {
			s.Name = "&sect;"
			s.Type = `<font color="yellow">Phenomena</font>`
		} else {
			s.Name = "&OElig;"
			s.Type = `<font color="magenta">Location</font>`
		}

		if sigNameLoc != "" {
			s.Description = sigNameLoc + `&nbsp;<span style="font-size: small;">(` + sigName + `)</span>`
		} else {
			s.Description = sigName
		}

		signals = append(signals, s)
	}

	// predict and add bodies BIO signals
	for name, bsigs := range cs.bios.Predict(cs) {
		s := &SignalT{
			Name:        utils.HTMLSafe(cs.BodyName(name)),
			Type:        `<font color="limegreen">Biological</font>`,
			Description: bsigs[0],
			Hints:       bsigs[1],
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
