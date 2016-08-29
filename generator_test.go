package muid

import (
	"encoding/binary"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGenerator(t *testing.T) {
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
}
func TestGenerate(t *testing.T) {
	machineid := []byte{0x66, 0x6f, 0x6f}
	var g *Generator
	var err error
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

func TestBulk(t *testing.T) {
	machineid := []byte{0x66, 0x6f, 0x6f}
	t.Run("successful bulk id generation", func(t *testing.T) {
		g, err := NewGenerator(3, 3, machineid)
		require.Nil(t, err)
		mm := g.Bulk(3)
		t.Run("correct number of MUIDs returned", func(t *testing.T) {
			assert.Equal(t, 3, len(mm))

		})
		t.Run("muids contain machineid bytes", func(t *testing.T) {
			assert.Equal(t, MUID(machineid), mm[0][3:6])
			assert.Equal(t, MUID(machineid), mm[1][3:6])
			assert.Equal(t, MUID(machineid), mm[2][3:6])
		})
		t.Run("muids are not identical", func(t *testing.T) {
			assert.NotEqual(t, mm[0], mm[1])
			assert.NotEqual(t, mm[1], mm[2])
			assert.NotEqual(t, mm[2], mm[0])
		})
		t.Run("muids contain sequential timestamps", func(t *testing.T) {
			g, err := NewGenerator(8, 3, machineid)
			require.Nil(t, err)

			mm := g.Bulk(3)
			t0 := binary.BigEndian.Uint64(mm[0][0:8])
			t1 := binary.BigEndian.Uint64(mm[1][0:8])
			t2 := binary.BigEndian.Uint64(mm[2][0:8])
			assert.Equal(t, t0+1, t1)
			assert.Equal(t, t1+1, t2)
		})
	})

	// Haven't yet figured out how to catch this panic with assert and avoid killing the tests:
	// t.Run("panics when simultaneous usage makes a collision", func(t *testing.T) {
	// 	assert.Panics(t, collide)
	// })
}

// func collide() {
// 	machineid := []byte{0x66, 0x6f, 0x6f}
// 	g, _ := NewGenerator(3, 3, machineid)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		g.Bulk(1000)
// 	}()
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		g.Bulk(1000)
// 	}()
// 	wg.Wait()

// }
