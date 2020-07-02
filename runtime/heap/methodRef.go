package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 11:39
 */

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(pool *ConstantPool, info *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.pool = pool
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
