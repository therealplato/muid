package muid

import (
	"encoding/binary"
	"fmt"
	"time"
)

const sizeBytes = 16                   // bytes
const sizeLeft = 8                     // bytes for timestamp
const sizeRight = sizeBytes - sizeLeft // bytes for machine id

// MUID is similar to a UUIDv1
// The MSB bits represent a timestamp, the LSB bits are a generator machine ID
type MUID []byte

// Generate creates a MUID
func Generate(machineID []byte) (MUID, error) {
	if machineID == nil {
		return nil, fmt.Errorf("missing %d required machineID bytes", sizeRight)
	}
	// TODO: Try starting with padOrTrim(machineID, sizeBytes)
	id := make([]byte, sizeBytes)
	machineID = padOrTrim(machineID, sizeRight)
	copy(id[sizeLeft:], machineID)
	t := time.Now().UnixNano()
	binary.BigEndian.PutUint64(id[:sizeLeft], uint64(t)) // thx http://stackoverflow.com/a/11015354/1380669
	return MUID(id), nil
}

// String returns a hex formatted id string
// Casting to []byte avoids a recursion with fmt.Sprintf
func (id *MUID) String() string {
	return fmt.Sprintf("%0x", []byte(*id))
}

// padOrTrim returns (size) bytes from input (bb)
// Short bb gets zeros prefixed, Long bb gets left/MSB bits trimmed
func padOrTrim(bb []byte, size int) []byte {
	l := len(bb)
	if l == size {
		return bb
	}
	if l > size {
		return bb[l-size:]
	}
	tmp := make([]byte, size)
	copy(tmp[size-l:], bb)
	return tmp
}
