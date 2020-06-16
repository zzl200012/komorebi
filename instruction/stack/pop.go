package stack

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 15:16
 */

/* POP could only pop vars like int and float that just take 1 operandStack size */
type POP struct { base.NoOperandsInstruction }
/* And when it comes to double or long, we use POP2 to pop them out */
type POP2 struct { base.NoOperandsInstruction }

func (instruction *POP) Execute(frame *runtime.Frame) {
	frame.OperandStack().PopSlot()
}
func (instruction *POP2) Execute(frame *runtime.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}