package runtime

import "komorebi/runtime/heap"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 20:29
 */

type Slot struct {
	num int32
	ref *heap.Object
}
