package edsm

import (
	"io/ioutil"
	"net/http"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

var cached map[string]string

func init() {
	cached = make(map[string]string)
}

// work in progress
func Query(system string) string {

	data := ""

	if found, ok := cached[system]; ok {
		return found
	}

	query := sconf.StrDef("EDSM", "systems url", "https://www.edsm.net/api-v1/system") +
		"?systemName=" + system + "&showInformation=1&showPrimaryStar=1"
	//"?systemName=" + url.PathEscape(system) + "&showInformation=1&showPrimaryStar=1"

	slog.Debug(5, "EDSM QUERY: %s", query)

	if resp, err := http.Get(query); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil && len(body) > len(system) { // a crutch
			data = string(body)
		}
	} else {
		slog.Err("EDSM returned: %s", err)
	}

	cached[system] = data

	return data
}
