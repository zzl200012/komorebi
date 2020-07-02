package runtime

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 16:29
 */

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // stack is implemented as linked list
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if stack._top != nil {
		frame.lower = stack._top
	}

	stack._top = frame
	stack.size++
}

func (stack *Stack) pop() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}

	top := stack._top
	stack._top = top.lower
	top.lower = nil
	stack.size--

	return top
}

func (stack *Stack) top() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}

	return stack._top
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}
