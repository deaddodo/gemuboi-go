package cpu

import "fmt"

type operation func(uint8, interface{}, interface{})

func (c *LR35902) opStop(op uint8, operandA interface{}, operandB interface{}) {
	c.Running = false
	c.Locked = true
}

func (c *LR35902) opHalt(op uint8, operandA interface{}, operandB interface{}) {
	c.Running = false
}

func (c *LR35902) opADD(op uint8, operandA interface{}, operandB interface{}) {
	if v, ok := operandA.(register); ok {
		fmt.Printf("%b\n", v.GetUint16())
	}
}

func (c *LR35902) opLD(op uint8, operandA interface{}, operandB interface{}) {
	// nope
}

// InstructionDecode contains the main execution logic of the CPU
func (c *LR35902) InstructionDecode() {
	c.Reg.AF.hi = 12

	for c.Running && !c.Locked {
		if c.Interrupted {
			// Execute interrupt
		} else {
			var operandA interface{}
			var operandB interface{}
			var opCall operation
			//var opLocation = c.Reg.PC

			//################# μOps #################
			//var opcode = c.consumeUint8()
			var opcode uint8 = 0x10

			/* Operand setup */
			// right hand operations are much more common, check those first
			if otWr[opcode] {
				//operandB = c.consumeUint8()
			} else if otDWr[opcode] {
				//operandB = c.consumeUint16()
			} else if otWl[opcode] {
				//operandA = c.consumeUint8()
			} else if otDWl[opcode] {
				//operandA = c.consumeUint16()
			}

			if otRegAl[opcode] {
				operandA = c.Reg.AF
			}

			if otRegAr[opcode] {
				operandB = c.Reg.AF
			}

			/* Cycle management */

			/* Operation */
			// only one operation can be selected. use a switch to short circuit
			// map checks
			switch {
			case opcode == 0x10:
				opCall = c.opStop
			case otLD[opcode]:
				opCall = c.opLD
			case otADD[opcode]:
				opCall = c.opADD
			}
			//################# μOps #################

			// execute the operation
			if opCall != nil {
				//fmt.Printf("%x [0x%x] <%#v> <%#v>\n", opLocation, opcode, operandA, operandB)

				opCall(opcode, operandA, operandB)
			} else {
				//fmt.Printf("! 0x%x\n", opcode)
			}
		}
	}
}
