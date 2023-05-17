package events

import (
	"errors"
	"reflect"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"

	"wedpad/internal/utils"
)

type EventHandler struct {
	cs     *CurrentSystemT
	bios   *BiosT
	sounds struct{} // TBD
}

func (evh *EventHandler) Init() error {

	loadedData := make(map[string][]byte)

	// load all data
	if err := utils.LoadDir(loadedData, sconf.StrDef("paths", "data", "data"), ".json", 100_000, 32); err != nil {
		return err
	}

	// bios data loaded - init its handler
	if bioData, ok := loadedData["bios"]; ok {
		evh.bios = new(BiosT)
		if err := evh.bios.Init(bioData); err != nil {
			return err
		}
	} else {
		return errors.New("No BIO data loaded (missing 'bios.json'?)")
	}

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
