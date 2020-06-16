package math

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 19:47
 */

/* Shift left int */
type ISHL struct{ base.NoOperandsInstruction }

/* there are only 32 bits for int, so we just need five bit of v2 to represent the offset */
func (instruction *ISHL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// the offset should be uint in Go
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

/* Arithmetic shift right int */
type ISHR struct{ base.NoOperandsInstruction }

func (instruction *ISHR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

/* Logical shift right int */
type IUSHR struct{ base.NoOperandsInstruction }

func (instruction *IUSHR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

/* Shift left long */
type LSHL struct{ base.NoOperandsInstruction }

func (instruction *LSHL) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

/* Arithmetic shift right long */
type LSHR struct{ base.NoOperandsInstruction }

func (instruction *LSHR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

/* Logical shift right long */
type LUSHR struct{ base.NoOperandsInstruction }

func (instruction *LUSHR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
