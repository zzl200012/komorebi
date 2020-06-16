package control

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 11:33
 */

// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (instruction *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	instruction.defaultOffset = reader.ReadInt32()
	instruction.low = reader.ReadInt32()
	instruction.high = reader.ReadInt32()
	jumpOffsetsCount := instruction.high - instruction.low + 1
	instruction.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (instruction *TABLE_SWITCH) Execute(frame *runtime.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= instruction.low && index <= instruction.high {
		offset = int(instruction.jumpOffsets[index-instruction.low])
	} else {
		offset = int(instruction.defaultOffset)
	}

	base.Branch(frame, offset)
}
