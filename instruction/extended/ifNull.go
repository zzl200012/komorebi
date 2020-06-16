package extended

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 14:54
 */

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

func (instruction *IFNULL) Execute(frame *runtime.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, instruction.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (instruction *IFNONNULL) Execute(frame *runtime.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, instruction.Offset)
	}
}
