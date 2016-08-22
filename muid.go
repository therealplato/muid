package muid

import "fmt"

const length = 16           // bytes
const left = 8              // bytes for timestamp
const right = length - left // bytes for machine id

// MUID is similar to a UUIDv1
// A portion of its bits represent a timestamp, the rest is the generator
// machine ID
type MUID []byte

// Generate creates a MUID
func Generate(machineID []byte) MUID {
	machineID = padOrTrim(machineID, right)
	// t := time.Now().UnixNano()
	// fmt.Printf("%016x\n", t)
	return MUID{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40, 0x00,
		0x80, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
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
