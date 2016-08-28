package main

import (
	"encoding/hex"
	"flag"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"

	"github.com/therealplato/muid"
)

func main() {
	f, _ := os.Create("cpuprofile.out")
	g, _ := os.Create("memprofile.out")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	var count = flag.Int("n", 1000, "generate this many ID's")
	// var quiet = flag.Bool("q", false, "mute stdout")
	var bytesT = flag.Int("sizets", 8, "use this many bytes for LSBs of nanosecond timestamp")
	var bytesMID = flag.Int("sizemid", 8, "use this many bytes for machine ID")
	var midhex = flag.String("mid", "1234567890abcdef", "hexadecimal machine id")
	flag.Parse()
	machineID, err := hex.DecodeString(*midhex)
	if err != nil {
		panic("machineID is invalid bytes: " + err.Error())
	}
	if len(machineID) != *bytesMID {
		panic("machineID is wrong number of hex bytes")
	}
	m, err := muid.NewGenerator(*bytesT, *bytesMID, machineID)
	if err != nil {
		panic(err)
	}
	// for i := 0; i < *count; i++ {
	// 	tmp := m.Generate()
	// 	out[i] = tmp
	// 	// if !*quiet {
	// 	// 	fmt.Printf("%x\n", tmp)
	// 	// }
	// }
	out := m.Bulk(*count)
	pprof.WriteHeapProfile(g)
	g.Close()
	_ = out
}
