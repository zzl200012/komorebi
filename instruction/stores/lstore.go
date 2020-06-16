package stores

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:42
 */

/* Store long into local variable */
type LSTORE struct{ base.Index8Instruction }

func lStore(frame *runtime.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

func (instruction *LSTORE) Execute(frame *runtime.Frame) {
	lStore(frame, instruction.Index)
}
func (instruction *LSTORE_0) Execute(frame *runtime.Frame) {
	lStore(frame, 0)
}
func (instruction *LSTORE_1) Execute(frame *runtime.Frame) {
	lStore(frame, 1)
}
func (instruction *LSTORE_2) Execute(frame *runtime.Frame) {
	lStore(frame, 2)
}
func (instruction *LSTORE_3) Execute(frame *runtime.Frame) {
	lStore(frame, 3)
}