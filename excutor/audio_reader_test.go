package excutor

import "testing"

func TestPlay(t *testing.T) {
	path := "./sample.mp3"
	err := Play(path)
	if err != nil {
		t.Error(err)
	}
}
