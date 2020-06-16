package math

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
	"math"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 15:35
 */

/* Remainder for double, float, int and long */
type DREM struct { base.NoOperandsInstruction }
type FREM struct { base.NoOperandsInstruction }
type IREM struct { base.NoOperandsInstruction }
type LREM struct { base.NoOperandsInstruction }

func (instruction *IREM) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: devide by zero")
	}
	result := val1 % val2
	stack.PushInt(result)
}

func (instruction *LREM) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: devide by zero")
	}
	result := val1 % val2
	stack.PushLong(result)
}

func (instruction *DREM) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	result := math.Mod(val1, val2)
	stack.PushDouble(result)
}

func (instruction *FREM) Execute(frame *runtime.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	result := math.Mod(float64(val1), float64(val2))
	stack.PushFloat(float32(result))
}

