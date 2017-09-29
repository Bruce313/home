package excutor

import (
	"os"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
)

func Play(filePath string) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	decoder, err := mp3.NewDecoder(f)
	if err != nil {
		return
	}
	defer decoder.Close()

	player, err := oto.NewPlayer(decoder.SampleRate(), 2, 2, 9182)
	if err != nil {
		return
	}
	defer player.Close()

	if _, err = io.Copy(player, decoder); err != nil {
		return
	}
	return nil
}