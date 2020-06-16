package conversions

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 20:18
 */

/* convert double type to * type */
type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (instruction *D2I) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (instruction *D2F) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

func (instruction *D2L) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
