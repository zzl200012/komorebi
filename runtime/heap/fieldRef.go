package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 11:39
 */

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(pool *ConstantPool, info *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.pool = pool
	ref.copyMemberRefInfo(&info.ConstantMemberRefInfo)
	return ref
}

func (ref *FieldRef) ResolvedField() *Field {
	if ref.field == nil {
		ref.resolveFieldRef()
	}
	return ref.field
}

// jvms 5.4.3.2
func (ref *FieldRef) resolveFieldRef() {
	d := ref.pool.class
	c := ref.ResolvedClass()
	field := lookupField(c, ref.name, ref.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	ref.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iF := range c.interfaces {
		if field := lookupField(iF, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
