package main

import "fmt"

type registers struct {
	AF, BC, DE, HL, SP, PC uint16
}

// LR35902 represents the primary CPU interface for the GB/GBC.
type LR35902 struct {
	Memory  *MemoryIO
	PPU     *DMGPPU
	Reg     registers
	Running bool
}

// Init initializes the LR35902 to it's cold-boot state
func (c *LR35902) Init(Memory *MemoryIO, PPU *DMGPPU) {
	c.Memory = Memory
	c.PPU = PPU
	c.Running = true
	c.Reg.PC = 0x0000
}

func (c *LR35902) consumeUint8(addr uint16) uint8 {
	c.Reg.PC++
	return c.Memory.GetUint8(addr)
}

func (c *LR35902) consumeUint16(addr uint16) uint16 {
	c.Reg.PC += 2
	return c.Memory.GetUint16(addr)
}

func (c *LR35902) instructionDecode() {
	for c.Running {
		//fmt.Printf("Opcode: 0x%x", c.Memory.GetUint8(c.Reg.PC))
		if c.Memory.GetUint8(c.Reg.PC) == 0x10 {
			// STOP opcode
			fmt.Println("STOP")
			c.Running = false
		}

		/******************************** μOp table *******************************/
		// Need to convert this from Rust gemuboi.
		/******************************** μOp table *******************************/

		c.Reg.PC++
	}
}

// Run is the primary execute cycle for the emulator
// TODO: probably want to have the CPU/PPU be separate goroutines
//       down the road and have them communicate through a channel
//       on sync
func (c *LR35902) Run() error {
	fmt.Println("Beginning instruction cycle...")
	c.instructionDecode()
	fmt.Println("Instructions exhausted.")

	return nil
}
