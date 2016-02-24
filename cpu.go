package gemuboi

import (
	"fmt"
)

// Registers represent the LR35902's basic registers
type Registers struct {
	AF, BC, DE, HL, SP, PC uint16
}

// LR35902 represents the primary CPU interface for the GB/GBC.
type LR35902 struct {
	Memory *MemoryIO
	Reg    Registers
}

// Init initializes the LR35902 to it's cold-boot state
func (c *LR35902) Init(Memory *MemoryIO) {
	c.Memory = Memory
	c.Reg.PC = 0x0000
}

// Run is the primary execute cycle for the emulator
// TODO: probably want to have the CPU/PPU be separate goroutines
//       down the road and have them communicate through a channel
//       on sync
func (c *LR35902) Run() error {
	fmt.Printf("Opcode: 0x%x\n", c.Memory.GetUint8(0x101))
	fmt.Printf("Value: 0x%x\n", c.Memory.GetUint16(0x102))

	fmt.Printf("MBC: 0x%x\n", c.Memory.GetUint8(0x147))

	return nil
}
