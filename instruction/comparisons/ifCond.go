package comparisons

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 10:56
 */

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.BranchInstruction }

func (instruction *IFEQ) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (instruction *IFNE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (instruction *IFLT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (instruction *IFLE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (instruction *IFGT) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, instruction.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (instruction *IFGE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, instruction.Offset)
	}
}
