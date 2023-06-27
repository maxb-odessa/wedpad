package events

import (
	"fmt"
	"sort"
	"strings"
	"wedpad/internal/msg"
	"wedpad/internal/utils"

	"github.com/fvbommel/sortorder"
)

func (cs *CurrentSystemT) ShowNotes() {

	notes := make([]map[string]interface{}, 0)

	planets := cs.PlanetsByName()

	var keys []string

	for k, _ := range planets {
		keys = append(keys, k)
	}

	sort.Strings(sortorder.Natural(keys))

	notesCnt := 0

	for _, name := range keys {

		body := planets[name]
		id := body.BodyID

		if nob := cs.notesOnBody(id); len(nob) > 0 {
			note := make(map[string]interface{})
			note["Notes"] = strings.Join(nob, "<br>")
			note["Body"] = utils.HTMLSafe(cs.BodyName(name))
			notes = append(notes, note)
			notesCnt++
		}

	}

	if notesCnt > 0 {

		m := &msg.Message{
			Target: msg.TARGET_NOTES,
			Type:   msg.TYPE_VIEW,
			Action: msg.ACTION_REPLACE,
			Data:   notes,
		}
		m.Send()

		buttonText := fmt.Sprintf("%d", notesCnt)
		m = &msg.Message{
			Target: msg.TARGET_NOTES,
			Type:   msg.TYPE_BUTTON,
			Action: msg.ACTION_REPLACE,
			Data:   buttonText,
		}
		m.Send()

		m = &msg.Message{
			Target: msg.TARGET_NOTES,
			Type:   msg.TYPE_BUTTON,
			Action: msg.ACTION_ATTENTION,
			Data:   "",
		}
		m.Send()
	}

}
