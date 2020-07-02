package reference

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
	"komorebi/runtime/heap"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 20:15
 */

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }

func (instruction *INSTANCE_OF) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(instruction.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
