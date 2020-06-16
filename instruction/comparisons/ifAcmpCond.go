package comparisons

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 11:08
 */

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }

func (instruction *IF_ACMPEQ) Execute(frame *runtime.Frame) {
	if aCmp(frame) {
		base.Branch(frame, instruction.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (instruction *IF_ACMPNE) Execute(frame *runtime.Frame) {
	if !aCmp(frame) {
		base.Branch(frame, instruction.Offset)
	}
}

func aCmp(frame *runtime.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
