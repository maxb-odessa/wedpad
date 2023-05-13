package events

import (
	"github.com/maxb-odessa/sconf"
)

func (cs *CurrentSystemT) IsRemarkableBody(id int) bool {

	if cs.gotSignals(id) {
		return true
	}

	if cs.wideRings(id) {
		return true
	}

	if cs.niceAtmosphere(id) {
		return true
	}

	return false
}

func (cs *CurrentSystemT) gotSignals(id int) bool {

	wantBio := int(sconf.Int32Def("criteria", "min bio signals", 0))
	wantGeo := int(sconf.Int32Def("criteria", "min geo signals", 0))
	wantHuman := int(sconf.Int32Def("criteria", "min human signals", 0))
	wantOther := int(sconf.Int32Def("criteria", "min other signals", 0))

	// no signals detected but at least one is wanted
	sigs, ok := cs.PlanetSignals()[id]
	if !ok {
		if wantBio+wantGeo+wantHuman+wantOther > 0 {
			return false
		} else {
			return true
		}
	}

	for _, sig := range sigs.Signals {
		switch sig.Type {
		case "$SAA_SignalType_Biological":
			if sig.Count >= wantBio {
				return true
			}
		case "$SAA_SignalType_Geological":
			if sig.Count >= wantGeo {
				return true
			}
		case "$SAA_SignalType_Human":
			if sig.Count >= wantHuman {
				return true
			}
		case "$SAA_SignalType_Other":
			if sig.Count >= wantOther {
				return true
			}
		}

	}

	return false
}

func (cs *CurrentSystemT) wideRings(id int) bool {
	return false
}

func (cs *CurrentSystemT) niceAtmosphere(id int) bool {
	return false
}
