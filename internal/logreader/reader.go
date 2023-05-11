package logreader

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"

	"github.com/danwakefield/fnmatch"
	"github.com/nxadm/tail"
	"github.com/radovskyb/watcher"
)

type reader struct {
	files   []string
	dir     string
	mask    string
	linesCh chan string
	pathCh  chan string
}

func (r *reader) run() error {

	// set files dir (default is current)
	dir := sconf.StrDef("journal", "dir", "./")
	r.dir, _ = filepath.Abs(os.ExpandEnv(dir))
	r.dir += "/"

	// set files mask (default is *.log)
	r.mask = sconf.StrDef("journal", "mask", `*.log`)

	r.linesCh = make(chan string, 64)
	r.pathCh = make(chan string, 1)

	// start dir watcher
	if err := r.watchDir(); err != nil {
		return err
	}

	// start tailer
	go r.tailFile()

	return nil
}

func (r *reader) read() string {

	select {
	case line, ok := <-r.linesCh:
		if ok {
			return line
		}
	}
	return ""
}

func (r *reader) getRecentFile() string {

	sortAsc := true

	if sconf.StrDef("journal", "sort", "asc") != "asc" {
		sortAsc = false
	}

	if len(r.files) > 0 {
		sort.Strings(r.files)
		if sortAsc {
			return r.files[len(r.files)-1]
		} else {
			return r.files[0]
		}
	}

	return ""
}

func (r *reader) watchDir() error {

	// monitor the direcotory for newer file to appear

	w := watcher.New()
	w.FilterOps(watcher.Create)

	filterFunc := func(info os.FileInfo, fullPath string) error {
		if fnmatch.Match(r.dir+r.mask, fullPath, 0) {
			return nil
		}
		return watcher.ErrSkip
	}

	w.AddFilterHook(filterFunc)

	if err := w.Add(r.dir); err != nil {
		return err
	}

	for path, st := range w.WatchedFiles() {
		if !st.IsDir() {
			r.files = append(r.files, path)
		}
	}

	if p := r.getRecentFile(); p != "" {
		r.pathCh <- p
	}

	// start dir watcher
	go func() {
		for {
			select {
			case event := <-w.Event:
				r.files = append(r.files, event.Path)
				if rf := r.getRecentFile(); rf != "" {
					r.pathCh <- r.getRecentFile()
				}
			case err := <-w.Error:
				slog.Err("file watcher error: %s", err)
			case <-w.Closed:
				return
			}
		}
	}()

	go w.Start(time.Second * 1)

	return nil
}

func (r *reader) tailFile() {

	ctx, cancel := context.WithCancel(context.Background())

	for {

		select {
		case path, ok := <-r.pathCh:
			slog.Debug(9, "file changed to: '%s'", path)
			if !ok {
				break
			}
			cancel()
			ctx, cancel = context.WithCancel(context.Background())
			go r.realTailer(ctx, path)
		} //select

	} //for

}

func (r *reader) realTailer(ctx context.Context, path string) {

	cfg := tail.Config{
		ReOpen: true,
		Follow: true,
		Poll:   true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
	}

	slog.Debug(5, "tailing '%s'", path)
	tailer, err := tail.TailFile(path, cfg)
	if err != nil {
		slog.Err("tailer failed on '%s': %s", path, err)
	}

	defer tailer.Stop()
	defer tailer.Cleanup()
	defer slog.Debug(5, "stopped tailing '%s'", path)

	for {
		select {
		case <-ctx.Done():
			return
		case line := <-tailer.Lines:
			r.linesCh <- line.Text
		}
	}
}
