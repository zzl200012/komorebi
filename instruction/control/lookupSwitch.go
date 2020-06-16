package control

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 11:36
 */

// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (instruction *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	instruction.defaultOffset = reader.ReadInt32()
	instruction.npairs = reader.ReadInt32()
	instruction.matchOffsets = reader.ReadInt32s(instruction.npairs * 2)
}

func (instruction *LOOKUP_SWITCH) Execute(frame *runtime.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < instruction.npairs*2; i += 2 {
		if instruction.matchOffsets[i] == key {
			offset := instruction.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(instruction.defaultOffset))
}
