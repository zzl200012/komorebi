package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 20:56
 */

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

type ConstantClassInfo struct {
	pool      ConstantPool
	nameIndex uint16
}

func (info *ConstantClassInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUint16()
}
func (info *ConstantClassInfo) Name() string {
	return info.pool.getUtf8(info.nameIndex)
}
