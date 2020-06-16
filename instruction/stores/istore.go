package stores

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:42
 */

/* Store int into local variable */
type ISTORE struct{ base.Index8Instruction }

func iStore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

type ISTORE_0 struct{ base.NoOperandsInstruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type ISTORE_3 struct{ base.NoOperandsInstruction }

func (instruction *ISTORE) Execute(frame *runtime.Frame) {
	iStore(frame, instruction.Index)
}
func (instruction *ISTORE_0) Execute(frame *runtime.Frame) {
	iStore(frame, 0)
}
func (instruction *ISTORE_1) Execute(frame *runtime.Frame) {
	iStore(frame, 1)
}
func (instruction *ISTORE_2) Execute(frame *runtime.Frame) {
	iStore(frame, 2)
}
func (instruction *ISTORE_3) Execute(frame *runtime.Frame) {
	iStore(frame, 3)
}