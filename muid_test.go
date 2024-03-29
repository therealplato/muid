package muid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMuidGenerate(t *testing.T) {
	t.Run("generate concatenates ts+mid", func(t *testing.T) {
		ts := []byte{0, 1, 2}
		machineid := []byte{0x66, 0x6f, 0x6f}
		id := generate(3, 3, ts, machineid)
		assert.Equal(t, len(id), 6)
		assert.Equal(t, MUID([]byte{0, 1, 2, 0x66, 0x6f, 0x6f}), id)
	})
}

func TestPadOrTrim(t *testing.T) {
	in := []byte{1, 2, 3, 4, 5}

	t.Run("given correct size input, returns same input", func(t *testing.T) {
		out := padOrTrim(in, 5)
		assert.Equal(t, []byte{1, 2, 3, 4, 5}, out)
	})

	t.Run("given long input, returns size bytes, rightmost/least significant bits", func(t *testing.T) {
		out := padOrTrim(in, 3)
		assert.Equal(t, []byte{3, 4, 5}, out)
	})

	t.Run("given short input, returns size bytes, zero prefixed", func(t *testing.T) {
		in := []byte{1, 2, 3, 4, 5}
		out := padOrTrim(in, 8)
		assert.Equal(t, []byte{0, 0, 0, 1, 2, 3, 4, 5}, out)
	})

	t.Run("given nil input, returns size bytes, zerod", func(t *testing.T) {
		in := []byte{}
		out := padOrTrim(in, 8)
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, out)
	})
}
