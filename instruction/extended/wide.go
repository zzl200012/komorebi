package extended

import (
	"komorebi/instruction/base"
	"komorebi/instruction/loads"
	"komorebi/instruction/math"
	"komorebi/instruction/stores"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 14:48
 */

/* Extend local variable index by additional bytes */
type WIDE struct {
	modifiedInstruction base.Instruction
}

func (instruction *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		instruction.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		instruction.modifiedInstruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (instruction *WIDE) Execute(frame *runtime.Frame) {
	instruction.modifiedInstruction.Execute(frame)
}
