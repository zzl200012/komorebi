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

// Check whether object is of given type
type CHECK_CAST struct{ base.Index16Instruction }

func (instruction *CHECK_CAST) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(instruction.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
