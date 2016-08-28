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
	SizeTS    int
	SizeMID   int
	MachineID []byte
}

// Generate generates one MUID
func (g *Generator) Generate() MUID {
	time.Sleep(1 * time.Nanosecond) // avoid generating multiple ID's within nanosecond timestamp resolution
	t := time.Now().UnixNano()
	ts := make([]byte, g.SizeTS)
	binary.BigEndian.PutUint64(ts, uint64(t)) // thx http://stackoverflow.com/a/11015354/1380669
	return generate(g.SizeTS, g.SizeMID, ts, g.MachineID)
}

// Bulk generates many MUIDs
func (g *Generator) Bulk(int) []MUID {
	return nil
}
