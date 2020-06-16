package base

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 13:46
 */

type Index8Instruction struct {
	Index uint
}

func (instruction *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	instruction.Index = uint(reader.ReadInt8())
}