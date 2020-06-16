package conversions

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 20:36
 */

/* Convert float to double */
type F2D struct{ base.NoOperandsInstruction }

func (instruction *F2D) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

/* Convert float to int */
type F2I struct{ base.NoOperandsInstruction }

func (instruction *F2I) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

/* Convert float to long */
type F2L struct{ base.NoOperandsInstruction }

func (instruction *F2L) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
