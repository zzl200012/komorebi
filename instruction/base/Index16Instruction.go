package base

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 13:47
 */

type Index16Instruction struct {
	Index uint
}

func (instruction *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	instruction.Index = uint(reader.ReadInt16())
}