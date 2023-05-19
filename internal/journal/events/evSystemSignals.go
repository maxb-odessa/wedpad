package events

import (
	"wedpad/internal/msg"

	"github.com/maxb-odessa/slog"
)

type SignalT struct {
	Name        string
	Type        string
	Description string
}

func (cs *CurrentSystemT) ShowSignals() {

	signals := make([]*SignalT, 0)

	// add system signals (phenomena, stations, etc)
	for _, sig := range cs.Signals() {

		s := new(SignalT)

		if sig.IsStation {
			s.Name = "?"
			s.Type = "Station"
			s.Description = sig.SignalName
		} else {
			// TODO for now: collect all sig.SignalName variants
			s.Name = "?"
			s.Type = "Phenomena"
			s.Description = sig.SignalNameLocalised + " (" + sig.SignalName + ")"
		}

		signals = append(signals, s)
	}

	// add planets predicted bios
	for _, sig := range cs.PlanetPredictedBioSignals() {
		signals = append(signals, sig)
	}

	sigsCnt := len(signals)
	slog.Debug(9, "SIGS CNT: %d", sigsCnt)
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
