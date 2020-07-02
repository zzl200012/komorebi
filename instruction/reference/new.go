package reference

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
	"komorebi/runtime/heap"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 20:16
 */

// Create new object
type NEW struct{ base.Index16Instruction }

func (instruction *NEW) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(instruction.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// todo: init class

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
