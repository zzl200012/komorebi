package comparisons

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 10:43
 */

/* Compare long type */
type LCMP struct{ base.NoOperandsInstruction }

func (instruction *LCMP) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
