package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 20:28
 */

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, infos []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(infos))
	for i, info := range infos {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(info)
		methods[i].copyAttributeInfo(info)
	}
	return methods
}

func (method *Method) copyAttributeInfo(info *classfile.MemberInfo) {
	if codeAttr := info.CodeAttribute(); codeAttr != nil {
		method.code = codeAttr.Code()
		method.maxStack = codeAttr.MaxStack()
		method.maxLocals = codeAttr.MaxLocals()
	}
}

func (method *Method) IsSynchronized() bool {
	return 0 != method.accessFlags&ACC_SYNCHRONIZED
}
func (method *Method) IsBridge() bool {
	return 0 != method.accessFlags&ACC_BRIDGE
}
func (method *Method) IsVarargs() bool {
	return 0 != method.accessFlags&ACC_VARARGS
}
func (method *Method) IsNative() bool {
	return 0 != method.accessFlags&ACC_NATIVE
}
func (method *Method) IsAbstract() bool {
	return 0 != method.accessFlags&ACC_ABSTRACT
}
func (method *Method) IsStrict() bool {
	return 0 != method.accessFlags&ACC_STRICT
}

// getters
func (method *Method) MaxStack() uint {
	return method.maxStack
}
func (method *Method) MaxLocals() uint {
	return method.maxLocals
}
func (method *Method) Code() []byte {
	return method.code
}
