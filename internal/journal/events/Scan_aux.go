package events

var Stars map[int]*ScanT

func cleanStars() {
	Stars = make(map[int]*ScanT)
}

var Planets map[int]*ScanT

func cleanPlanets() {
	Planets = make(map[int]*ScanT)
}

var BaryCentres map[int]*ScanBaryCentreT

func cleanBaryCentres() {
	BaryCentres = make(map[int]*ScanBaryCentreT)
}
