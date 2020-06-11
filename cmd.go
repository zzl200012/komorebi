package main

import (
	"flag"
	"fmt"
	"komorebi/classpath"
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
	XjreOption string
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
	fmt.Printf("classpath:%s class:%s args:%v\n",
		       cmd.CpOption, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Class Data: %v\n", classData)
}