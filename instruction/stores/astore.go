package stores

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 14:42
 */

/* Store reference into local variable */
type ASTORE struct { base.Index8Instruction }

func aStore(frame *runtime.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

type ASTORE_0 struct { base.NoOperandsInstruction }
type ASTORE_1 struct { base.NoOperandsInstruction }
type ASTORE_2 struct { base.NoOperandsInstruction }
type ASTORE_3 struct { base.NoOperandsInstruction }

func (instruction *ASTORE) Execute(frame *runtime.Frame) {
	aStore(frame, instruction.Index)
}
func (instruction *ASTORE_0) Execute(frame *runtime.Frame) {
	aStore(frame, 0)
}
func (instruction *ASTORE_1) Execute(frame *runtime.Frame) {
	aStore(frame, 1)
}
func (instruction *ASTORE_2) Execute(frame *runtime.Frame) {
	aStore(frame, 2)
}
func (instruction *ASTORE_3) Execute(frame *runtime.Frame) {
	aStore(frame, 3)
}


