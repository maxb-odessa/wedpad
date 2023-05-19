package events

import (
	"reflect"

	"github.com/maxb-odessa/slog"
)

type EventHandler struct {
	cs *CurrentSystemT
}

func (evh *EventHandler) Init() error {
	evh.cs = new(CurrentSystemT)
	return evh.cs.Init()
}

func (evh *EventHandler) CurrentSystem() *CurrentSystemT {
	return evh.cs
}

func (evh *EventHandler) Handle(evName string, evData map[string]interface{}) {
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(evData)

	val := reflect.ValueOf(evh).MethodByName(evName)
	if val.IsValid() {
		val.Call(in)
	} else {
		slog.Warn("Usupported journal event '%s', not handling it", evName)
	}
}
