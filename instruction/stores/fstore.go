package stores

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:42
 */

/* Store float into local variable */
type FSTORE struct{ base.Index8Instruction }

func fStore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

func (instruction *FSTORE) Execute(frame *runtime.Frame) {
	fStore(frame, instruction.Index)
}
func (instruction *FSTORE_0) Execute(frame *runtime.Frame) {
	fStore(frame, 0)
}
func (instruction *FSTORE_1) Execute(frame *runtime.Frame) {
	fStore(frame, 1)
}
func (instruction *FSTORE_2) Execute(frame *runtime.Frame) {
	fStore(frame, 2)
}
func (instruction *FSTORE_3) Execute(frame *runtime.Frame) {
	fStore(frame, 3)
}


