package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 14:25
 */

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(pool *ConstantPool,
	info *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.pool = pool
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}
