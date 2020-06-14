package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 10:52
 */

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (attr *UnparsedAttribute) readInfo(reader *ClassReader) {
	attr.info = reader.readBytes(attr.length)
}

func (attr *UnparsedAttribute) Info() []byte {
	return attr.info
}
