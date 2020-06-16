package base

import "komorebi/runtime"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 10:41
 */

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *runtime.Frame)
}