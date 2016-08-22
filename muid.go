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
func Generate() MUID {
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
