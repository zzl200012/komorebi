package comparisons

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 10:46
 */

/* Compare float type */
type FCMPL struct{ base.NoOperandsInstruction }
type FCMPG struct{ base.NoOperandsInstruction }

/* if any of the floats is NaN, fcmpg returns 1, but fcmpl returns -1 */
func fCmp(frame *runtime.Frame, flag bool) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else if val1 < val2 {
		stack.PushInt(-1)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (instruction *FCMPG) Execute(frame *runtime.Frame) {
	fCmp(frame, true)
}
func (instruction *FCMPL) Execute(frame *runtime.Frame) {
	fCmp(frame, false)
}
