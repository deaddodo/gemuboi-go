package main

import "fmt"

func (c *LR35902) opStop() {
	c.Running = false
	c.Locked = true
}

func (c *LR35902) opHalt() {
	c.Running = false
}

type operation func(uint8, interface{}, interface{})

func (c *LR35902) opAdd(op uint8, operandA interface{}, operandB interface{}) {
	// do nothing for now
}

func (c *LR35902) opLD(op uint8, operandA interface{}, operandB interface{}) {
	fmt.Println("LD call:")
	fmt.Printf("   0x%x operand: %#v/ operand: %#v\n", op, operandA, operandB)
}

// InstructionDecode contains the main execution logic of the CPU
func (c *LR35902) InstructionDecode() {
	for c.Running && !c.Locked {
		if c.Interrupted {
			// Execute interrupt
		} else {
			var operandA interface{}
			var operandB interface{}
			var opCall operation

			//################# μOps #################
			var opcode = c.consumeUint8()

			if otW[opcode] {
				operandA = c.consumeUint8()
			}
			if otDW[opcode] {
				operandA = c.consumeUint16()
			}
			if otLD[opcode] {
				opCall = c.opLD
			}
			//################# μOps #################

			// execute the operation
			if opcode == 0x10 {
				c.opStop()
			} else if opCall != nil {
				opCall(opcode, operandA, operandB)
			} else {
				fmt.Println("Opcode (0x%x) unimplemented!", opcode)
			}
		}
	}
}
