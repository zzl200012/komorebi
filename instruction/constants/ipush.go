package constants

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:07
 */

/* get a byte and extend it to int, then push it to the top */
type BIPUSH struct{ val int8 }

/* get a short ... ... */
type SIPUSH struct{ val int16 }

func (instruction *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	instruction.val = reader.ReadInt8()
}
func (instruction *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	instruction.val = reader.ReadInt16()
}
func (instruction *BIPUSH) Execute(frame *runtime.Frame) {
	i := int32(instruction.val)
	frame.OperandStack().PushInt(i)
}
func (instruction *SIPUSH) Execute(frame *runtime.Frame) {
	i := int32(instruction.val)
	frame.OperandStack().PushInt(i)
}
