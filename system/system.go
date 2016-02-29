package system

import (
	"gemuboi-go/cpu"
	"gemuboi-go/memory"
	"gemuboi-go/ppu"
)

// Bus provides the main bus through which the CPU/PPU/APU/Memory communicate
type Bus struct {
	mem memory.IO
	cpu cpu.LR35902
	ppu ppu.DMGPPU
}

// Init bootstraps the inital system configuration
func (sys *Bus) Init(biosFile string, romFile string) {
	sys.mem.Init(biosFile, romFile)
	sys.cpu.Init()
	sys.ppu.Init()
}

// PowerOn starts the system execution
func (sys *Bus) PowerOn() {
	sys.cpu.Start()
}
