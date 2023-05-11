package events

import (
	"github.com/maxb-odessa/slog"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

func (cs *CurrentSystemT) ShowPlanets() {

	// those are for BODIES view port
	bodies := make(map[string]interface{})
	for _, ev := range cs.planets {

		if !cs.remarkable("planet", ev) {
			continue
		}

	}

	// those are for LOG view port
	for _, ev := range cs.planets {

		if !cs.remarkable("body", ev) {
			continue
		}

	}

}

type luaData struct {
	ev  *ScanT
	cs  *CurrentSystemT
	res bool
}

func (d *luaData) SetResult(r bool) {
	d.res = r
}

func (cs *CurrentSystemT) remarkable(what string, ev *ScanT) bool {

	scriptName := "remarkable-" + what

	script, err := cs.FindLua(scriptName)
	if err != nil {
		slog.Err("LUA: Failed to call '%s.lua': %s", scriptName, err)
		return false
	}

	l := lua.NewState()
	defer l.Close()

	ld := &luaData{
		ev:  ev,
		cs:  cs,
		res: false,
	}

	l.SetGlobal("data", luar.New(l, ld))
	if err := l.DoString(script); err != nil {
		slog.Err("LUA: Failed to call '%s.lua': %s", scriptName, err)
		return false
	}

	slog.Debug(5, "LUA: '%s' returned '%+v'", scriptName, ld.res)

	return ld.res
}
