package main

import (
	"log"
	"os"
)

func parseArgs() string {
	if len(os.Args) != 2 {
		log.Fatal("File name necessary for ROM load")
	} else {
		if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			log.Fatal("File provided does not exist")
		}
	}

	return os.Args[1]
}

func main() {
	var Memory MemoryIO
	var CPU LR35902
	var PPU DMGPPU

	if Memory.LoadROM(parseArgs()) != nil {
		log.Fatal("ROM data failed to load")
	}

	CPU.Init(&Memory, &PPU)
	CPU.Run()
}
