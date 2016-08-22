package main

import (
	"fmt"

	"github.com/therealplato/muid"
)

func main() {
	// Generate one ID:
	i := muid.Generate()
	fmt.Println(i.String())
}
