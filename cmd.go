package main

import (
	"flag"
	"fmt"
	"komorebi/classfile"
	"komorebi/classpath"
	"komorebi/runtime"
	"os"
	"strings"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 10:09
 */

// java [-options] class [args...]
type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
	XjreOption  string
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}

func StartJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	//fmt.Printf("classpath:%s class:%s args:%v\n",
	//	       cmd.CpOption, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	//fmt.Println(cmd.Class)
	//printClassInfo(cf)
	//frame := runtime.NewFrame(100, 100)
	//testLocalVars(frame.LocalVars())
	//testOperandStack(frame.OperandStack())
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.Class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}

func testLocalVars(vars runtime.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *runtime.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
