package math

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 20:00
 */

// Negate double
type DNEG struct{ base.NoOperandsInstruction }

func (instruction *DNEG) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOperandsInstruction }

func (instruction *FNEG) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOperandsInstruction }

func (instruction *INEG) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOperandsInstruction }

func (instruction *LNEG) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
