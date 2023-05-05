package msg

import (
	"bytes"
	"errors"
	"html/template"
	"strings"

	"wedpad/internal/server"
	"wedpad/internal/utils"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

type Message struct {
	Action string
	Target string
	Type   string
	Data   interface{}
}

const (
	TARGET_TOP     = "top"
	TARGET_SYSTEM  = "system"
	TARGET_LOG     = "log"
	TARGET_BODIES  = "bodies"
	TARGET_SIGNALS = "signals"
	TARGET_BOTTOM  = "bottom"

	TYPE_VIEW   = "view"
	TYPE_BUTTON = "button"

	ACTION_NONE      = "none"
	ACTION_CLEAN     = "clean"
	ACTION_APPEND    = "append"
	ACTION_REPLACE   = "replace"
	ACTION_ATTENTION = "attention"
)

var templates map[string]*template.Template

func Init() error {

	templates = make(map[string]*template.Template)

	files := make(map[string][]byte)

	dir := sconf.StrDef("paths", "templates", "templates")

	if err := utils.LoadDir(files, dir, ".tmpl", 8192, 32); err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("No HTML templates loaded")
	}

	for n, t := range files {
		if tmpl, err := template.New(n).Parse(string(t)); err != nil {
			return err
		} else {
			templates[n] = tmpl
			slog.Debug(9, "Added temlpate '%s'", n)
		}
	}

	server.SetOnConnect(showGreeting)

	return nil
}

func showGreeting() {

	data := make(map[string]string)
	data["text"] = "WEdPad - Web based Elite Dangerous Pad"

	msg := &Message{
		Action: ACTION_REPLACE,
		Target: TARGET_LOG,
		Type:   TYPE_VIEW,
		Data:   data,
	}

	msg.Send()

}

type serverMsg struct {
	Action string
	Target string
	Type   string
	Data   template.HTML
}

func (msg *Message) Send() error {

	sm := serverMsg{
		Action: msg.Action,
		Target: msg.Target,
		Type:   msg.Type,
		Data:   "",
	}

	target := msg.Type + "-" + msg.Target

	if tmpl, ok := templates[target]; ok {
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, msg.Data); err != nil {
			slog.Warn("Template '%s' execution failed: %s", target, err)
		} else {
			res := strings.ReplaceAll(buf.String(), "\n", "")
			sm.Data = template.HTML(res)
			slog.Debug(9, "String after templating: '%s'", sm.Data)
		}
	} else {
		slog.Warn("Template '%s' is not defined", target)
	}

	if d, err := utils.JSONMarshal(sm); err != nil {
		slog.Err("JSON Marshall failed: %s", err)
		return err
	} else {
		server.Send(d)
	}

	return nil
}
