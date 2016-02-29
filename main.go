package main

import (
	"flag"
	"gemuboi-go/system"
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
	var System system.Bus

	System.Init(parseArgs())
	System.PowerOn()
}
