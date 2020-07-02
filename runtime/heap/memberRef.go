package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 11:34
 */

type MemberRef struct {
	Ref
	name       string
	descriptor string
}

func (ref *MemberRef) copyMemberRefInfo(info *classfile.ConstantMemberRefInfo) {
	ref.className = info.ClassName()
	ref.name, ref.descriptor = info.NameAndDescriptor()
}

func (ref *MemberRef) Name() string {
	return ref.name
}
func (ref *MemberRef) Descriptor() string {
	return ref.descriptor
}
