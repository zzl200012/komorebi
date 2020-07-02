package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 20:24
 */

/* Both of field and method are the members of class, and they have some same fields, so here we define a
struct ClassMember */
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (member *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	member.accessFlags = memberInfo.AccessFlags()
	member.name = memberInfo.Name()
	member.descriptor = memberInfo.Descriptor()
}

/* to judge if this member is XXX or not */
func (member *ClassMember) IsProtected() bool {
	return member.accessFlags&ACC_PROTECTED != 0
}
func (member *ClassMember) IsPublic() bool {
	return member.accessFlags&ACC_PUBLIC != 0
}
func (member *ClassMember) IsPrivate() bool {
	return member.accessFlags&ACC_PRIVATE != 0
}
func (member *ClassMember) IsFinal() bool {
	return member.accessFlags&ACC_FINAL != 0
}
func (member *ClassMember) IsStatic() bool {
	return member.accessFlags&ACC_STATIC != 0
}
func (member *ClassMember) IsSynthetic() bool {
	return member.accessFlags&ACC_SYNTHETIC != 0
}

/* getters */
func (member *ClassMember) Name() string {
	return member.name
}
func (member *ClassMember) Descriptor() string {
	return member.descriptor
}
func (member *ClassMember) Class() *Class {
	return member.class
}

func (member *ClassMember) isAccessibleTo(d *Class) bool {
	if member.IsPublic() {
		return true
	}
	c := member.class
	if member.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !member.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
