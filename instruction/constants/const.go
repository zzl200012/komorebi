package constants

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 13:55
 */

/* The following instructions just push the const inside the operands to the top of OperandStack */

// Push null
type ACONST_NULL struct { base.NoOperandsInstruction }

func (instruction *ACONST_NULL) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushRef(nil)
}

// Push double
type DCONST_0 struct { base.NoOperandsInstruction }

func (instruction *DCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct { base.NoOperandsInstruction }

func (instruction *DCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Push float
type FCONST_0 struct { base.NoOperandsInstruction }

func (instruction *FCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct { base.NoOperandsInstruction }

func (instruction *FCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct { base.NoOperandsInstruction }

func (instruction *FCONST_2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Push int constant
type ICONST_M1 struct { base.NoOperandsInstruction }

func (instruction *ICONST_M1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct { base.NoOperandsInstruction }

func (instruction *ICONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct { base.NoOperandsInstruction }

func (instruction *ICONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct { base.NoOperandsInstruction }

func (instruction *ICONST_2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct { base.NoOperandsInstruction }

func (instruction *ICONST_3) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct { base.NoOperandsInstruction }

func (instruction *ICONST_4) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct { base.NoOperandsInstruction }

func (instruction *ICONST_5) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(5)
}

// Push long constant
type LCONST_0 struct { base.NoOperandsInstruction }

func (instruction *LCONST_0) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct { base.NoOperandsInstruction }

func (instruction *LCONST_1) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushLong(1)
}
