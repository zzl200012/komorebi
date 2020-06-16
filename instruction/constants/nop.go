package constants

import (
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/15 13:52
 */

/* NOP declares doing nothing */
type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Execute(frame *runtime.Frame) {}
