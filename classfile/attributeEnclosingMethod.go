package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 10:49
 */

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	pool        ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (attr *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	attr.classIndex = reader.readUint16()
	attr.methodIndex = reader.readUint16()
}

func (attr *EnclosingMethodAttribute) ClassName() string {
	return attr.pool.getClassName(attr.classIndex)
}

func (attr *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if attr.methodIndex > 0 {
		return attr.pool.getNameAndType(attr.methodIndex)
	} else {
		return "", ""
	}
}
