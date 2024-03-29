package muid

import (
	"encoding/binary"
	"errors"
	"time"
)

// NewGenerator takes the byte count for the timestamp bits, the byte count for
// the machine ID, and a byte slice machine ID. It will generate unique ID's
// given that the system clock is set correctly and no other generator is
// using the same machine ID simultaneously
func NewGenerator(sizeTS, sizeMID int, mid []byte) (*Generator, error) {
	if sizeTS < 1 {
		return nil, errors.New("sizeTS must be at least one")
	}
	if sizeMID < 1 {
		return nil, errors.New("sizeMID must be at least one")
	}
	if len(mid) != sizeMID {
		return nil, errors.New("missing required machineID bytes")
	}

	return &Generator{
		SizeTS:    sizeTS,
		SizeMID:   sizeMID,
		MachineID: mid,
	}, nil
}

// Generator generates MUIDs
type Generator struct {
	SizeTS    int // bytes of timestamp
	SizeMID   int // bytes of machine id
	MachineID []byte
	LastTS    uint64 // unix nanoseconds
}

// Generate generates one MUID based on the current system time
func (g *Generator) Generate() MUID {
	time.Sleep(1 * time.Nanosecond) // avoid generating multiple ID's within nanosecond timestamp resolution
	t := time.Now().UnixNano()
	ts := make([]byte, 8)
	binary.BigEndian.PutUint64(ts, uint64(t)) // thx http://stackoverflow.com/a/11015354/1380669
	return generate(g.SizeTS, g.SizeMID, padOrTrim(ts, g.SizeTS), g.MachineID)
}

// Bulk generates many MUIDs with sequential timestamps, i.e. if Bulk(3) is
// called at timestamp 100, the ids will have timestamps 100, 101, 102
func (g *Generator) Bulk(n int) []MUID {
	if n < 1 {
		panic("need >0 bulk generator count")
	}
	results := make([]MUID, n)
	t0 := uint64(time.Now().UnixNano())
	if t0 < g.LastTS {
		panic("race condition, generated bulk ids too fast")
	}
	g.LastTS = t0 + uint64(n)
	for i := uint64(0); i < uint64(n); i++ {
		ts := make([]byte, 8)
		binary.BigEndian.PutUint64(ts, uint64(t0+i)) // thx http://stackoverflow.com/a/11015354/1380669
		results[i] = generate(g.SizeTS, g.SizeMID, padOrTrim(ts, g.SizeTS), g.MachineID)
	}
	return results
}
