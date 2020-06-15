package runtime

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 16:27
 */

/*
 * JVM -> Pc
 *     -> Stack -> Frame -> Local Variable
 *                       -> Operand Stack
 */

type Thread struct {
	pc int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (thread *Thread) PC() int {
	return thread.pc
}
func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}

func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}
func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top()
}
