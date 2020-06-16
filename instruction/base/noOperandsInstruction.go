package base

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 10:49
 */

/* NoOperandsInstruction declares there is no opCode, so nothing to do */
type NoOperandsInstruction struct {}
func (instruction *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {}