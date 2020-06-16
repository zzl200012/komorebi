package loads

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:15
 */

/* Load float from local variable */
type FLOAD struct{ base.Index8Instruction }

func fLoad(frame *runtime.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

func (instruction *FLOAD) Execute(frame *runtime.Frame) {
	fLoad(frame, instruction.Index)
}
func (instruction *FLOAD_0) Execute(frame *runtime.Frame) {
	fLoad(frame, 0)
}
func (instruction *FLOAD_1) Execute(frame *runtime.Frame) {
	fLoad(frame, 1)
}
func (instruction *FLOAD_2) Execute(frame *runtime.Frame) {
	fLoad(frame, 2)
}
func (instruction *FLOAD_3) Execute(frame *runtime.Frame) {
	fLoad(frame, 3)
}


