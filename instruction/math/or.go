package math

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 19:59
 */

/* Boolean OR int */
type IOR struct{ base.NoOperandsInstruction }

func (instruction *IOR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

/* Boolean OR long */
type LOR struct{ base.NoOperandsInstruction }

func (instruction *LOR) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
