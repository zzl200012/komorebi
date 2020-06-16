package runtime

import "math"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 20:29
 */

type OperandStack struct {
	size uint
	slots []Slot
}

func newOperandStack(maxSize uint) *OperandStack {
	if maxSize <= 0 {
		return nil
	}
	return &OperandStack{
		slots: make([]Slot, maxSize),
	}
}
func (operandStack *OperandStack) PushInt(val int32) {
	operandStack.slots[operandStack.size].num = val
	operandStack.size++
}

func (operandStack *OperandStack) PopInt() int32 {
	operandStack.size--
	return operandStack.slots[operandStack.size].num
}

func (operandStack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	operandStack.slots[operandStack.size].num = int32(bits)
	operandStack.size++
}
func (operandStack *OperandStack) PopFloat() float32 {
	operandStack.size--
	bits := uint32(operandStack.slots[operandStack.size].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (operandStack *OperandStack) PushLong(val int64) {
	operandStack.slots[operandStack.size].num = int32(val)
	operandStack.slots[operandStack.size+1].num = int32(val >> 32)
	operandStack.size += 2
}
func (operandStack *OperandStack) PopLong() int64 {
	operandStack.size -= 2
	low := uint32(operandStack.slots[operandStack.size].num)
	high := uint32(operandStack.slots[operandStack.size+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (operandStack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	operandStack.PushLong(int64(bits))
}
func (operandStack *OperandStack) PopDouble() float64 {
	bits := uint64(operandStack.PopLong())
	return math.Float64frombits(bits)
}

func (operandStack *OperandStack) PushRef(ref *Object) {
	operandStack.slots[operandStack.size].ref = ref
	operandStack.size++
}
func (operandStack *OperandStack) PopRef() *Object {
	operandStack.size--
	ref := operandStack.slots[operandStack.size].ref
	operandStack.slots[operandStack.size].ref = nil
	return ref
}
func (operandStack *OperandStack) PushSlot(slot Slot) {
	operandStack.slots[operandStack.size] = slot
	operandStack.size++
}
func (operandStack *OperandStack) PopSlot() Slot {
	operandStack.size--
	return operandStack.slots[operandStack.size]
}

