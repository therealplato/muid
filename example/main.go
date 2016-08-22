package main

import (
	"fmt"

	"github.com/therealplato/muid"
)

func main() {
	machineID := []byte("foo")
	// Generate one ID:
	i := muid.Generate(machineID)
	fmt.Println(i.String())
}
