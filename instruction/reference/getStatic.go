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

// Get static field from class
type GET_STATIC struct{ base.Index16Instruction }

func (instruction *GET_STATIC) Execute(frame *runtime.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(instruction.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	// todo: init class

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
