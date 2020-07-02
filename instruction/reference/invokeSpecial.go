package reference

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 20:15
 */

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (instruction *INVOKE_SPECIAL) Execute(frame *runtime.Frame) {
	frame.OperandStack().PopRef()
}
