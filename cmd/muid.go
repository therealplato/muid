package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"

	"github.com/therealplato/muid"
)

func main() {
	var count = flag.Int("n", 1000, "generate this many ID's")
	var quiet = flag.Bool("q", false, "mute stdout")
	var bytesT = flag.Int("sizets", 8, "use this many bytes for LSBs of nanosecond timestamp")
	var bytesMID = flag.Int("sizemid", 8, "use this many bytes for machine ID")
	var midhex = flag.String("mid", "1234567890abcdef", "hexadecimal machine id")
	var profile = flag.Bool("profile", false, "enable cpu and memory profiling")
	flag.Parse()

	var f, g *os.File
	if *profile {
		f, _ = os.Create("cpuprofile.out")
		g, _ = os.Create("memprofile.out")
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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

	out := m.Bulk(*count)
	if !*quiet {
		for _, muid := range out {
			fmt.Printf("%x\n", muid)
		}
	}
	if *profile {
		pprof.WriteHeapProfile(g)
		g.Close()
	}
}
