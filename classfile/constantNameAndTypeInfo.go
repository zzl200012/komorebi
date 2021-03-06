package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 20:52
 */

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (info *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUint16()
	info.descriptorIndex = reader.readUint16()
}