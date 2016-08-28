package muid

// NewGenerator takes an integer size in bytes for timestamp bits, and a byte
// slice machine ID. It will generate unique ID's given that the system clock
// is set correctly and no other generator is using the same machine ID
// simultaneously
func NewGenerator(sizeTS, sizeMID int, mid []byte) (*Generator, error) {
	return nil, nil
}

// Generator generates MUIDs
type Generator struct {
	SizeTS    int
	SizeMID   int
	MachineID []byte
}

// Generate generates one MUID
func (g *Generator) Generate() MUID {
	muid, _ := Generate(g.SizeTS+g.SizeMID, g.SizeTS, g.MachineID)
	return muid
}

// Bulk generates many MUIDs
func (g *Generator) Bulk(int) []MUID {
	return nil
}
