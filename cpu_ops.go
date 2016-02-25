package main

func (c *LR35902) opSTOP() {
	c.Running = false
}

// InstructionDecode contains the main execution logic of the CPU
func (c *LR35902) InstructionDecode() {
	for c.Running {
		/******************************** μOp table *******************************/
		switch c.consumeUint8() {
		case 0x10:
			c.opSTOP()
		}
		/******************************** μOp table *******************************/
	}
}
