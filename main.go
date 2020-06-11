package main

import (
	"fmt"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/11 10:09
 */

func main() {
	c := ParseCmd()
	if c.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if c.HelpFlag || c.Class == "" {
		PrintUsage()
	} else {
		StartJVM(c)
	}
}
