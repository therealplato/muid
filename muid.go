package muid

import (
	"encoding/binary"
	"fmt"
	"time"
)

// MUID is similar to a UUIDv1
// The MSB bits represent a timestamp, the LSB bits are a generator machineID
type MUID []byte

// Generate takes the total id length sizeBytes, the number of bytes in the
// timestamp portion sizeLeft, and a machineID. If machineID is empty, an error
// is returned. Otherwise, machineID is left zero padded to (sizeBytes-sizeLeft)
// and then truncated to (sizeBytes-sizeLeft) bytes. A MUID is constructed by
// concatenating the timestamp bytes and machineID bytes and returned.
func generate(sizeBytes, sizeLeft int, machineID []byte) MUID {
	sizeRight := sizeBytes - sizeLeft // bytes for machine id
	// TODO: Try starting with padOrTrim(machineID, sizeBytes)
	id := make([]byte, sizeBytes)
	machineID = padOrTrim(machineID, sizeRight)
	copy(id[sizeLeft:], machineID)
	t := time.Now().UnixNano()
	binary.BigEndian.PutUint64(id[:sizeLeft], uint64(t)) // thx http://stackoverflow.com/a/11015354/1380669
	return MUID(id)
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
