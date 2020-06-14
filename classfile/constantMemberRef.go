package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 21:18
 */

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldRefInfo struct{ ConstantMemberRefInfo }
type ConstantMethodRefInfo struct{ ConstantMemberRefInfo }
type ConstantInterfaceMethodRefInfo struct{ ConstantMemberRefInfo }

type ConstantMemberRefInfo struct {
	pool             ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (info *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	info.classIndex = reader.readUint16()
	info.nameAndTypeIndex = reader.readUint16()
}

func (info *ConstantMemberRefInfo) ClassName() string {
	return info.pool.getClassName(info.classIndex)
}
func (info *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return info.pool.getNameAndType(info.nameAndTypeIndex)
}
