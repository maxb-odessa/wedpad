package events

import (
	"fmt"
	"sort"
	"strings"
	"wedpad/internal/msg"
	"wedpad/internal/utils"
)

func (cs *CurrentSystemT) ShowNotes() {

	notes := make([]map[string]interface{}, 0)

	planets := cs.Planets()

	var keys []int

	for k, _ := range planets {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	notesCnt := 0

	for _, id := range keys {

		note := make(map[string]interface{})

		body := planets[id]

		if nob := cs.notesOnBody(id); len(nob) > 0 {
			note["Notes"] = strings.Join(nob, "<br>")
			note["Body"] = utils.HTMLSafe(cs.BodyName(body.BodyName))
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
