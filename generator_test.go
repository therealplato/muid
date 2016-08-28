package muid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	machineid := []byte{0x66, 0x6f, 0x6f}
	var g *Generator
	var err error
	t.Run("constructing a Generator", func(t *testing.T) {
		g, err = NewGenerator(3, 3, machineid)
		assert.Nil(t, err)
		t.Run("when machineid is blank", func(t *testing.T) {
			g, err = NewGenerator(3, 0, []byte{})
			assert.NotNil(t, err)
			assert.Nil(t, g)
		})
	})
	// 	t.Run("successful id generation", func(t *testing.T) {

	// 		t.Run("id is the correct length", func(t *testing.T) {
	// 			id, err := generate(, 8, machineid)
	// 			assert.Nil(t, err)
	// 			assert.Equal(t, len(id), 16)
	// 		})

	// 		t.Run("timestamp portion of id is recent", func(t *testing.T) {
	// 			id, err := generate(16, 8, machineid)
	// 			require.Nil(t, err)
	// 			tsb := id[:8]
	// 			tsi := binary.BigEndian.Uint64(tsb)
	// 			ts := time.Unix(0, int64(tsi))
	// 			assert.WithinDuration(t, time.Now(), ts, time.Millisecond)
	// 		})

	// 		t.Run("machineid portion of id", func(t *testing.T) {
	// 			t.Run("is left zero padded when short", func(t *testing.T) {
	// 				id, err := generate(16, 8, machineid)
	// 				assert.Nil(t, err)
	// 				expected := MUID{0, 0, 0, 0, 0, 0x66, 0x6f, 0x6f}
	// 				assert.Equal(t, expected, id[8:])
	// 			})
	// 			t.Run("is truncated, keeping right bytes, when long", func(t *testing.T) {
	// 				machineid := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 				id, err := generate(16, 8, machineid)
	// 				assert.Nil(t, err)
	// 				expected := MUID([]byte{
	// 					2, 3, 4, 5, 6, 7, 8, 9,
	// 				})
	// 				assert.Equal(t, expected, id[8:])
	// 			})
	// 		})
	// 		t.Run("repeated runs yield unique ids", func(t *testing.T) {
	// 			id1, _ := generate(16, 8, machineid)
	// 			id2, _ := generate(16, 8, machineid)
	// 			id3, _ := generate(16, 8, machineid)
	// 			assert.NotEqual(t, id1, id2)
	// 			assert.NotEqual(t, id2, id3)
	// 			assert.NotEqual(t, id3, id1)
	// 		})
	// 	})
	// 	t.Run("it returns an error when called with a blank machine id", func(t *testing.T) {
	// 		id, err := generate(16, 8, []byte{})
	// 		assert.Nil(t, id)
	// 		assert.NotNil(t, err)
	// 	})
	// }

	// func TestPadOrTrim(t *testing.T) {
	// 	in := []byte{1, 2, 3, 4, 5}

	// 	t.Run("given correct size input, returns same input", func(t *testing.T) {
	// 		out := padOrTrim(in, 5)
	// 		assert.Equal(t, []byte{1, 2, 3, 4, 5}, out)
	// 	})

	// 	t.Run("given long input, returns size bytes from the right", func(t *testing.T) {
	// 		out := padOrTrim(in, 3)
	// 		assert.Equal(t, []byte{3, 4, 5}, out)
	// 	})

	// 	t.Run("given short input, returns size bytes, zero prefixed", func(t *testing.T) {
	// 		in := []byte{1, 2, 3, 4, 5}
	// 		out := padOrTrim(in, 8)
	// 		assert.Equal(t, []byte{0, 0, 0, 1, 2, 3, 4, 5}, out)
	// 	})

	// 	t.Run("given nil input, returns size bytes, zero prefixed", func(t *testing.T) {
	// 		in := []byte{}
	// 		out := padOrTrim(in, 8)
	// 		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, out)
	// 	})
}
