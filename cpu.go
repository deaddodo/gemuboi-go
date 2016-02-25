package main

import "fmt"

type registers struct {
	AF, BC, DE, HL, SP, PC uint16
}

type interrupt struct {
	// STUB
}

// LR35902 represents the primary CPU interface for the GB/GBC.
type LR35902 struct {
	Memory                       *MemoryIO
	PPU                          *DMGPPU
	Reg                          registers
	Interrupt                    interrupt
	Running, Interrupted, Locked bool
}

// Init initializes the LR35902 to it's cold-boot state
func (c *LR35902) Init(Memory *MemoryIO, PPU *DMGPPU) {
	c.Memory = Memory
	c.PPU = PPU

	c.softReset()
}

func (c *LR35902) softReset() {
	c.Running = true
	c.Reg.PC = 0x0000
}

func (c *LR35902) consumeUint8() uint8 {
	var rVal = c.Memory.GetUint8(c.Reg.PC)
	c.Reg.PC++
	return rVal
}

func (c *LR35902) consumeUint16() uint16 {
	var rVal = c.Memory.GetUint16(c.Reg.PC)
	c.Reg.PC += 2
	return rVal
}

// Start spins up the execution cycle and deals with CPU/PPU sync
// TODO: probably want to have the CPU/PPU be separate goroutines
//       down the road and have them communicate through a channel
//       on sync
func (c *LR35902) Start() error {
	fmt.Println("Beginning instruction cycle...")
	c.InstructionDecode()
	fmt.Println("Instructions exhausted.")

	return nil
}
