package muid

import "fmt"

// MUID is similar to a UUIDv1
// The MSB bits represent a timestamp, the LSB bits are a generator machineID
type MUID []byte

// generate takes the timestamp byte count sizeTS, machine id byte count
// sizeMID, and the actual byte slices of each. mid is concatenated onto ts
// and returned. The function panics if a wrong length is specified.
func generate(sizeTS, sizeMID int, ts, mid []byte) MUID {
	if len(ts) != sizeTS || len(mid) != sizeMID {
		panic("muid: generate received bytes not matching given lengths")
	}
	size := sizeTS + sizeMID
	// left pad mid with zeroes up to total size:
	id := padOrTrim(mid, size)
	copy(ts, id[:sizeTS])
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
