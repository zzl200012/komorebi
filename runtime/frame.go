package runtime

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 16:31
 */

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters & setters
func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}
func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}
func (frame *Frame) Thread() *Thread {
	return frame.thread
}
func (frame *Frame) NextPC() int {
	return frame.nextPC
}
func (frame *Frame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}
