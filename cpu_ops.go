package main

func (c *LR35902) opStop() {
	c.Running = false
	c.Locked = true
}

// InstructionDecode contains the main execution logic of the CPU
func (c *LR35902) InstructionDecode() {
	for c.Running && !c.Locked {
		if c.Interrupted {
			// Execute interrupt
		} else {
			switch c.consumeUint8() {
			case 0x10:
				c.opStop()
			}
		}
	}
}
