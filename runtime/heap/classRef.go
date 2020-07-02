package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 11:31
 */

type ClassRef struct {
	Ref
}

func newClassRef(pool *ConstantPool, info *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.pool = pool
	ref.className = info.Name()
	return ref
}
