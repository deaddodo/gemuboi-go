package gob

import (
	"fmt"
)

type Registers struct {
	AF, BC, DE, HL, SP, PC uint16
}

type LR35902 struct {
	Memory *MemoryIO
	Reg    Registers
}

func (c *LR35902) Init(Memory *MemoryIO) {
	c.Memory = Memory
	c.Reg.PC = 0x0000
}

func (c *LR35902) Run() error {
	fmt.Printf("Opcode: 0x%x\n", c.Memory.GetUint8(0x101))
	fmt.Printf("Value: 0x%x\n", c.Memory.GetUint16(0x102))

	fmt.Printf("MBC: 0x%x\n", c.Memory.GetUint8(0x147))

	return nil
}
