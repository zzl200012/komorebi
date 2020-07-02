package heap

import "komorebi/classfile"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 20:28
 */

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, infos []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(infos))
	for i, info := range infos {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(info)
	}
	return fields
}

func (field *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		field.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (field *Field) IsVolatile() bool {
	return 0 != field.accessFlags&ACC_VOLATILE
}
func (field *Field) IsTransient() bool {
	return 0 != field.accessFlags&ACC_TRANSIENT
}
func (field *Field) IsEnum() bool {
	return 0 != field.accessFlags&ACC_ENUM
}

func (field *Field) ConstValueIndex() uint {
	return field.constValueIndex
}
func (field *Field) SlotId() uint {
	return field.slotId
}
func (field *Field) isLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}
