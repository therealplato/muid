package muid

import (
	"encoding/binary"
	"fmt"
	"time"
)

const length = 16           // bytes
const left = 8              // bytes for timestamp
const right = length - left // bytes for machine id

// MUID is similar to a UUIDv1
// A portion of its bits represent a timestamp, the rest is the generator
// machine ID
type MUID []byte

// Generate creates a MUID
func Generate(machineID []byte) MUID {
	id := make([]byte, length)
	machineID = padOrTrim(machineID, right)
	copy(id[left:], machineID)
	t := time.Now().UnixNano()
	binary.BigEndian.PutUint64(id[:left], uint64(t)) // thx http://stackoverflow.com/a/11015354/1380669
	return MUID(id)
}

// String returns a hex formatted id string
// Casting to []byte avoids a recursion with fmt.Sprintf
func (id *MUID) String() string {
	return fmt.Sprintf("%0x", []byte(*id))
}

// padOrTrim returns size bytes from input bb
// Short bb gets zeros prefixed
// Long bb gets left/MSB bits trimmed
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
