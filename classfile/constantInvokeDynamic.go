package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 21:56
 */

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (info *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	info.referenceKind = reader.readUint8()
	info.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (info *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	info.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (info *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	info.bootstrapMethodAttrIndex = reader.readUint16()
	info.nameAndTypeIndex = reader.readUint16()
}
