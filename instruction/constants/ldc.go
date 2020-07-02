package constants

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

// Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }

func (instruction *LDC) Execute(frame *runtime.Frame) {
	_ldc(frame, instruction.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (instruction *LDC_W) Execute(frame *runtime.Frame) {
	_ldc(frame, instruction.Index)
}

func _ldc(frame *runtime.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (instruction *LDC2_W) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(instruction.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}
