package comparisons

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 10:54
 */

/* Compare double type */
type DCMPL struct{ base.NoOperandsInstruction }
type DCMPG struct{ base.NoOperandsInstruction }

/* if any of the doubles is NaN, dcmpg returns 1, but dcmpl returns -1 */
func dCmp(frame *runtime.Frame, flag bool) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
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

func (instruction *DCMPG) Execute(frame *runtime.Frame) {
	dCmp(frame, true)
}
func (instruction *DCMPL) Execute(frame *runtime.Frame) {
	dCmp(frame, false)
}
