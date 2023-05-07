package events

import (
	"reflect"

	"github.com/maxb-odessa/slog"
)

type EventHandler struct{}

func (evHandler EventHandler) Init() {
	cleanStars()
	cleanPlanets()
	cleanBaryCentres()
}

func (evHandler EventHandler) Handle(evName string, evData map[string]interface{}) {
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(evData)

	val := reflect.ValueOf(&evHandler).MethodByName(evName)
	if val.IsValid() {
		val.Call(in)
	} else {
		slog.Err("Usupported journal event '%s', not handling it", evName)
	}
}
