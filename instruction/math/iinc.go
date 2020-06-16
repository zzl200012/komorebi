package math

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 20:12
 */

/* Increment local variable by constant */
type IINC struct {
	Index uint
	Const int32
}

func (instruction *IINC) FetchOperands(reader *base.BytecodeReader) {
	instruction.Index = uint(reader.ReadUint8())
	instruction.Const = int32(reader.ReadInt8())
}

func (instruction *IINC) Execute(frame *runtime.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(instruction.Index)
	val += instruction.Const
	localVars.SetInt(instruction.Index, val)
}
