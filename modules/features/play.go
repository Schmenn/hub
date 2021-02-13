package features

import (
	"os"
	"strings"
	"time"


	"github.com/Schmenn/hub/modules/debug"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

var(
	streamer beep.StreamSeekCloser
	format beep.Format
)

// Play plays song
func Play(s string) {
	ext := strings.Split(s, ".")

	f, err := os.Open(s)
	debug.Check(err)

	switch strings.ToLower(ext[len(ext)-1]) {
	case "wav":
		streamer, format, err = wav.Decode(f)
	case "mp3":
		streamer, format, err = mp3.Decode(f)
	case "ogg":
		streamer, format, err = vorbis.Decode(f)
	case "flac":
		streamer, format, err = flac.Decode(f)
	}
	debug.Check(err)

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
	select {}
}
