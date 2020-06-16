package loads

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:15
 */

/* Load double from local variable */
type DLOAD struct{ base.Index8Instruction }

func dLoad(frame *runtime.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }

func (instruction *DLOAD) Execute(frame *runtime.Frame) {
	dLoad(frame, instruction.Index)
}
func (instruction *DLOAD_0) Execute(frame *runtime.Frame) {
	dLoad(frame, 0)
}
func (instruction *DLOAD_1) Execute(frame *runtime.Frame) {
	dLoad(frame, 1)
}
func (instruction *DLOAD_2) Execute(frame *runtime.Frame) {
	dLoad(frame, 2)
}
func (instruction *DLOAD_3) Execute(frame *runtime.Frame) {
	dLoad(frame, 3)
}
