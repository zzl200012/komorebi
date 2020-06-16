package loads

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:15
 */

/* Load int from local vars */
type ILOAD struct{ base.Index8Instruction }

type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }

func iLoad(frame *runtime.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

/* index for i_load came from operands */
func (instruction *ILOAD) Execute(frame *runtime.Frame) {
	iLoad(frame, instruction.Index)
}

/* Another load instructions' index came from their opCode */
func (instruction *ILOAD_0) Execute(frame *runtime.Frame) {
	iLoad(frame, 0)
}
func (instruction *ILOAD_1) Execute(frame *runtime.Frame) {
	iLoad(frame, 1)
}
func (instruction *ILOAD_2) Execute(frame *runtime.Frame) {
	iLoad(frame, 2)
}
func (instruction *ILOAD_3) Execute(frame *runtime.Frame) {
	iLoad(frame, 3)
}
