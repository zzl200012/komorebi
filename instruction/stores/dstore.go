package stores

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:42
 */

/* Store double into local variable */
type DSTORE struct { base.Index8Instruction }

func dStore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

type DSTORE_0 struct { base.NoOperandsInstruction }
type DSTORE_1 struct { base.NoOperandsInstruction }
type DSTORE_2 struct { base.NoOperandsInstruction }
type DSTORE_3 struct { base.NoOperandsInstruction }

func (instruction *DSTORE) Execute(frame *runtime.Frame) {
	dStore(frame, instruction.Index)
}
func (instruction *DSTORE_0) Execute(frame *runtime.Frame) {
	dStore(frame, 0)
}
func (instruction *DSTORE_1) Execute(frame *runtime.Frame) {
	dStore(frame, 1)
}
func (instruction *DSTORE_2) Execute(frame *runtime.Frame) {
	dStore(frame, 2)
}
func (instruction *DSTORE_3) Execute(frame *runtime.Frame) {
	dStore(frame, 3)
}


