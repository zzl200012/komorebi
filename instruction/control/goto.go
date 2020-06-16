package control

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 11:28
 */

/* just jump without conditions */
type GOTO struct{ base.BranchInstruction }

func (instruction *GOTO) Execute(frame *runtime.Frame) {
	base.Branch(frame, instruction.Offset)
}
