package vm

const MAXARG_Bx = 1<<18 - 1
const MAXARG_sBx = MAXARG_Bx >> 1

type Instruction uint32

/*
 * 前6字节用于标识指令
 */
func (self Instruction) Opcode() int {
	return int(self & 0x3f)
}

/*
 * ABC, ABx, AsBx, Ax
 * 用于标识指令的参数个数, 及其所占的位数
 * ABC:  8+9+9
 * ABx:  8+18
 * AsBx: 8+18
 * Ax:   26
 */
func (self Instruction) ABC() (a, b, c int) {
	a = int(self >> 6 & 0xff)
	c = int(self >> 14 & 0x1ff)
	b = int(self >> 23 & 0x1ff)
	return
}

func (self Instruction) ABx() (a, b int) {
	a = int(self >> 6 & 0xff)
	b = int(self >> 14)
	return
}

func (self Instruction) Ax() (a int) {
	a = int(self >> 6)
	return
}

func (self Instruction) AsBx() (a, b int) {
	a, b = self.ABx()
	return a, b - MAXARG_sBx
}

func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}
