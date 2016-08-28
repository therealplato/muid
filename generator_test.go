package muid

import (
	"encoding/binary"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	machineid := []byte{0x66, 0x6f, 0x6f}
	var g *Generator
	var err error
	t.Run("constructing a Generator", func(t *testing.T) {
		g, err = NewGenerator(3, 3, machineid)
		assert.Nil(t, err)
		assert.Equal(t, 3, g.SizeTS)
		assert.Equal(t, 3, g.SizeMID)
		assert.Equal(t, machineid, g.MachineID)
		t.Run("errors when machineid is blank", func(t *testing.T) {
			g, err = NewGenerator(3, 0, []byte{})
			assert.NotNil(t, err)
			assert.Nil(t, g)
		})
		t.Run("errors when machineid is wrong length", func(t *testing.T) {
			g, err = NewGenerator(3, 2, []byte{1})
			assert.NotNil(t, err)
			assert.Nil(t, g)
		})
	})
	t.Run("successful id generation", func(t *testing.T) {
		g, err = NewGenerator(3, 3, machineid)
		id := g.Generate()

		t.Run("id is the correct length", func(t *testing.T) {
			assert.Nil(t, err)
			assert.Equal(t, len(id), 6)
		})
	})

	t.Run("timestamp portion of id is recent", func(t *testing.T) {
		g, err = NewGenerator(8, 3, machineid)
		require.Nil(t, err)
		id := g.Generate()
		tsb := id[:8]
		tsi := binary.BigEndian.Uint64(tsb)
		ts := time.Unix(0, int64(tsi))
		assert.WithinDuration(t, time.Now(), ts, time.Millisecond)
	})

	t.Run("repeated runs yield unique ids", func(t *testing.T) {
		g, err = NewGenerator(3, 3, machineid)
		id1 := g.Generate()
		id2 := g.Generate()
		id3 := g.Generate()
		assert.NotEqual(t, id1, id2)
		assert.NotEqual(t, id2, id3)
		assert.NotEqual(t, id3, id1)
	})
}
