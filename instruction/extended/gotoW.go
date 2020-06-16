package extended

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 14:48
 */

/* just jump without conditions ( with wide index ) */
type GOTO_W struct {
	offset int
}

func (instruction *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	instruction.offset = int(reader.ReadInt32())
}
func (instruction *GOTO_W) Execute(frame *runtime.Frame) {
	base.Branch(frame, instruction.offset)
}
