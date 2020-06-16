package base

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 11:06
 */

/* BranchInstruction declares 'jump', and FetchOperands reads a uint16 param, and converts it to int, then
assigns it to Offset */
type BranchInstruction struct {
	Offset int
}
func (instruction *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	instruction.Offset = int(reader.ReadInt16())
}
