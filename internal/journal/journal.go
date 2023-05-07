package journal

import (
	"encoding/json"
	"wedpad/internal/journal/events"
	"wedpad/internal/logreader"

	"github.com/maxb-odessa/slog"
)

func Init() error {
	go run()
	return nil
}

func run() {

	for {

		// wait for the log line
		line := logreader.GetLine()

		slog.Debug(9, "got line: '%s'", line)

		if len(line) == 0 {
			continue
		}

		var mData map[string]interface{}

		if err := json.Unmarshal([]byte(line), &mData); err != nil {
			slog.Warn("Journal entry json unmarshal failed: %s", err)
			continue
		}

		evh := new(events.EventHandler)

		if e, ok := mData["event"]; ok {
			evh.Handle(e.(string), mData)
		}
	}
}
