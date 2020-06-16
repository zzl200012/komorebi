package loads

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:15
 */

/* Load long from local variable */
type LLOAD struct{ base.Index8Instruction }

func lLoad(frame *runtime.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

func (instruction *LLOAD) Execute(frame *runtime.Frame) {
	lLoad(frame, instruction.Index)
}
func (instruction *LLOAD_0) Execute(frame *runtime.Frame) {
	lLoad(frame, 0)
}
func (instruction *LLOAD_1) Execute(frame *runtime.Frame) {
	lLoad(frame, 1)
}
func (instruction *LLOAD_2) Execute(frame *runtime.Frame) {
	lLoad(frame, 2)
}
func (instruction *LLOAD_3) Execute(frame *runtime.Frame) {
	lLoad(frame, 3)
}

