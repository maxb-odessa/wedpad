package utils

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/danwakefield/fnmatch"
	"github.com/maxb-odessa/slog"
)

// load no more than "maxNum"  files with extension "ext" of max size "maxSize" from dir "dir"
// and put them into map "holder" using filename without extension
func LoadDir(holder map[string][]byte, dir string, ext string, maxSize int64, maxNum int) error {

	absDir, _ := filepath.Abs(os.ExpandEnv(dir))

	d, err := os.Open(absDir)
	if err != nil {
		return err
	}

	defer d.Close()

	files, err := d.Readdir(0)
	if err != nil {
		return err
	}

	for _, f := range files {

		if len(holder) >= maxNum {
			break
		}

		if !f.Mode().IsRegular() {
			continue
		}

		if f.Size() > maxSize {
			continue
		}

		if !fnmatch.Match(`*`+ext, f.Name(), fnmatch.FNM_PATHNAME) {
			continue
		}

		path := absDir + `/` + f.Name()
		if data, err := os.ReadFile(path); err != nil {
			slog.Warn("Failed to read file '%s': %s", path, err)
		} else {
			slog.Debug(1, "Loaded file '%s'", path)
			noExt := strings.TrimSuffix(f.Name(), ext)
			holder[noExt] = data
		}

	}

	return nil
}

// marshal JSON but prevent HTML tags escaping
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
