package loads

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:15
 */

/* Load reference from local variable */
type ALOAD struct{ base.Index8Instruction }

func aLoad(frame *runtime.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
type ALOAD_0 struct{ base.NoOperandsInstruction }
type ALOAD_1 struct{ base.NoOperandsInstruction }
type ALOAD_2 struct{ base.NoOperandsInstruction }
type ALOAD_3 struct{ base.NoOperandsInstruction }

func (instruction *ALOAD) Execute(frame *runtime.Frame) {
	aLoad(frame, instruction.Index)
}
func (instruction *ALOAD_0) Execute(frame *runtime.Frame) {
	aLoad(frame, 0)
}
func (instruction *ALOAD_1) Execute(frame *runtime.Frame) {
	aLoad(frame, 1)
}
func (instruction *ALOAD_2) Execute(frame *runtime.Frame) {
	aLoad(frame, 2)
}
func (instruction *ALOAD_3) Execute(frame *runtime.Frame) {
	aLoad(frame, 3)
}


