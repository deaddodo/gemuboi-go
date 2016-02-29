package cpu

import (
	"fmt"
	//"gemuboi-go/system"
)

type register struct {
	hi, lo uint8
}

func (r *register) GetUint16() uint16 {
	return (uint16(r.hi) << 8) | uint16(r.lo)
}

func (r *register) SetUint16(value uint16) {
	r.lo = uint8(value)
	r.hi = uint8(value >> 8)
}

type registers struct {
	AF, BC, DE, HL register
	SP, PC         uint16
}

type interrupt struct {
	// STUB
}

// LR35902 represents the primary CPU interface for the GB/GBC.
type LR35902 struct {
	//sys                          *system.Bus
	Reg                          registers
	Interrupt                    interrupt
	Running, Interrupted, Locked bool
}

// Init initializes the LR35902 to it's cold-boot state
func (c *LR35902) Init( /*sys *system.Bus*/ ) {
	//c.sys = sys
	c.softReset()
}

func (c *LR35902) softReset() {
	c.Running = true
	c.Reg.PC = 0x0000
}

/*func (c *LR35902) consumeUint8() uint8 {
	var rVal = c.sys.mem.GetUint8(c.Reg.PC)
	c.Reg.PC++
	return rVal
}

func (c *LR35902) consumeUint16() uint16 {
	var rVal = c.sys.mem.GetUint16(c.Reg.PC)
	c.Reg.PC += 2
	return rVal
}*/

// Start spins up the execution cycle and deals with CPU/PPU sync
// TODO: probably want to have the CPU/PPU be separate goroutines
//       down the road and have them communicate through a channel
//       on sync
func (c *LR35902) Start() error {
	fmt.Println("Beginning instruction cycle...")
	c.InstructionDecode()
	fmt.Println("Instructions cycle complete.")

	return nil
}
