package comparisons

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 11:05
 */
// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }

func (instruction *IF_ICMPEQ) Execute(frame *runtime.Frame) {
	if val1, val2 := iCmpPop(frame); val1 == val2 {
		base.Branch(frame, instruction.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (instruction *IF_ICMPNE) Execute(frame *runtime.Frame) {
	if val1, val2 := iCmpPop(frame); val1 != val2 {
		base.Branch(frame, instruction.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (instruction *IF_ICMPLT) Execute(frame *runtime.Frame) {
	if val1, val2 := iCmpPop(frame); val1 < val2 {
		base.Branch(frame, instruction.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (instruction *IF_ICMPLE) Execute(frame *runtime.Frame) {
	if val1, val2 := iCmpPop(frame); val1 <= val2 {
		base.Branch(frame, instruction.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (instruction *IF_ICMPGT) Execute(frame *runtime.Frame) {
	if val1, val2 := iCmpPop(frame); val1 > val2 {
		base.Branch(frame, instruction.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (instruction *IF_ICMPGE) Execute(frame *runtime.Frame) {
	if val1, val2 := iCmpPop(frame); val1 >= val2 {
		base.Branch(frame, instruction.Offset)
	}
}

func iCmpPop(frame *runtime.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
