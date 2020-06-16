package base

import "komorebi/runtime"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 10:58
 */

func Branch(frame *runtime.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
