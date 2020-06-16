package main

import (
	"fmt"
	"komorebi/classfile"
	instruction2 "komorebi/instruction"
	"komorebi/instruction/base"
	"komorebi/runtime"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 14:55
 */

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttribute := methodInfo.CodeAttribute()
	maxLocals := codeAttribute.MaxLocals()
	maxStack := codeAttribute.MaxStack()
	thread := runtime.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catch(frame)
	loop(thread, codeAttribute.Code())
}

func loop(thread *runtime.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		reader.Reset(bytecode, pc)
		opCode := reader.ReadUint8()
		instruction := instruction2.NewInstruction(opCode)
		instruction.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, instruction, instruction)
		instruction.Execute(frame)
	}
}

func catch(frame *runtime.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
