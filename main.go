package main

import (
	"flag"
	"log"
	"os"
)

func parseArgs() (string, string) {
	biosFile := flag.String("bios", "DMG_ROM.bin", "location of the GB BIOS file")
	romFile := flag.String("rom", "file.gb", "location of the game ROM")
	flag.Parse()

	if _, err := os.Stat(*biosFile); os.IsNotExist(err) {
		log.Fatal("BIOS provided does not exist")
	}

	if _, err := os.Stat(*romFile); os.IsNotExist(err) {
		log.Fatal("ROM provided does not exist")
	}

	return *biosFile, *romFile
}

func main() {
	var Memory MemoryIO
	var CPU LR35902
	var PPU DMGPPU

	Memory.Init(parseArgs())
	CPU.Init(&Memory, &PPU)
	CPU.Start()
}
