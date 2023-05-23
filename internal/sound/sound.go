package sound

import (
	"bytes"
	"sync"
	"time"
	"wedpad/internal/utils"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

type track struct {
	sync.Mutex
	data   []byte
	buffer *beep.Buffer
}

type SoundT struct {
	queue  chan *track
	tracks map[string]*track
}

func (s *SoundT) Init() error {

	s.queue = make(chan *track, 16)
	s.tracks = make(map[string]*track)

	trData := make(map[string][]byte)

	if err := utils.LoadDir(trData, sconf.StrDef("paths", "sounds", "sounds"), ".wav", 1_000_000, 10); err != nil {
		return err
	}

	for trName, trWav := range trData {

		reader := bytes.NewReader(trWav)

		if streamer, format, err := wav.Decode(reader); err != nil {
			return err
		} else {
			tr := new(track)
			tr.buffer = beep.NewBuffer(format)
			tr.buffer.Append(streamer)
			streamer.Close()
			s.tracks[trName] = tr
			slog.Debug(1, "Loaded sound track '%s'", trName)
		}

	}

	var playSampleRate beep.SampleRate = 44100

	speaker.Init(playSampleRate, playSampleRate.N(time.Second/20))

	go s.run()

	return nil
}

func (s *SoundT) Play(track string) {
	if tr, ok := s.tracks[track]; ok {
		s.queue <- tr
	} else {
		slog.Warn("Failed to play '%s': not loaded?", track)
	}
}

func (s *SoundT) run() {
	for {
		select {
		case track := <-s.queue:
			track.play()
		}
	}
}

func (tr *track) play() {

	tr.Lock()
	defer tr.Unlock()

	stream := tr.buffer.Streamer(0, tr.buffer.Len())
	speaker.Play(stream)
}
