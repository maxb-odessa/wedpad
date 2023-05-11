package events

func (cs *CurrentSystemT) ShowPlanets() {

	// those are for BODIES view
	for _, ev := range cs.planets {

		if !interestingPlanet(ev) {
			continue
		}

	}

	// those are for LOG view

}

func interestingPlanet(ev *ScanT) bool {

	return true
}
