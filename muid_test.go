package muid

import (
	"encoding/binary"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("successful id generation", func(t *testing.T) {
		machineid := []byte{0x66, 0x6f, 0x6f}

		t.Run("id is the correct length", func(t *testing.T) {
			id, err := Generate(machineid)
			assert.Nil(t, err)
			assert.Equal(t, len(id), sizeBytes)
		})

		t.Run("timestamp portion of id is recent", func(t *testing.T) {
			require.True(t, sizeLeft >= 8) // otherwise the nanosecond timestamp gets left truncated
			id, err := Generate(machineid)
			require.Nil(t, err)
			tsb := id[:sizeLeft]
			tsi := binary.BigEndian.Uint64(tsb)
			ts := time.Unix(0, int64(tsi))
			assert.WithinDuration(t, time.Now(), ts, time.Millisecond)
		})

		t.Run("machineid portion of id", func(t *testing.T) {
			t.Run("is left zero padded when short", func(t *testing.T) {
				require.True(t, len(machineid) < sizeRight)
				id, err := Generate(machineid)
				assert.Nil(t, err)
				zeroes := sizeRight - len(machineid)
				zeroed := make([]byte, zeroes)
				zeroed = append(zeroed, 0x66, 0x6f, 0x6f)
				expected := MUID(zeroed)
				assert.Equal(t, expected, id[sizeLeft:])
			})
			t.Run("is truncated, keeping right bytes, when long", func(t *testing.T) {
				machineid := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
				require.True(t, len(machineid) > sizeRight)
				id, err := Generate(machineid)
				assert.Nil(t, err)
				expected := MUID([]byte{
					2, 3, 4, 5, 6, 7, 8, 9,
				})
				assert.Equal(t, expected, id[sizeLeft:])
			})
		})
	})
	t.Run("it returns an error when called with a blank machine id", func(t *testing.T) {
		id, err := Generate([]byte{})
		assert.Nil(t, id)
		assert.NotNil(t, err)
	})
}

func TestPadOrTrim(t *testing.T) {
	in := []byte{1, 2, 3, 4, 5}

	t.Run("given correct size input, returns same input", func(t *testing.T) {
		out := padOrTrim(in, 5)
		assert.Equal(t, []byte{1, 2, 3, 4, 5}, out)
	})

	t.Run("given long input, returns size bytes from the right", func(t *testing.T) {
		out := padOrTrim(in, 3)
		assert.Equal(t, []byte{3, 4, 5}, out)
	})

	t.Run("given short input, returns size bytes, zero prefixed", func(t *testing.T) {
		in := []byte{1, 2, 3, 4, 5}
		out := padOrTrim(in, 8)
		assert.Equal(t, []byte{0, 0, 0, 1, 2, 3, 4, 5}, out)
	})

	t.Run("given nil input, returns size bytes, zero prefixed", func(t *testing.T) {
		in := []byte{}
		out := padOrTrim(in, 8)
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, out)
	})
}