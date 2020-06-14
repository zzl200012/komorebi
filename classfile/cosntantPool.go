package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 15:52
 */

type ConstantPool []ConstantInfo

/* readConstantPool read the constant pool out */
func readConstantPool(reader *ClassReader) ConstantPool {
	count := int(reader.readUint16())
	pool := make([]ConstantInfo, count)
	for i := 1; i < count; i++ {
		pool[i] = readConstantInfo(reader, pool)
		switch pool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return pool
}
/* getConstantInfo seek for constant by index */
func (pool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info := pool[index]; info != nil {
		return info
	}
	panic("Invalid constant pool index!")
}
/* getNameAndType seek for name and descriptor of field or method from constant pool */
func (pool ConstantPool) getNameAndType(index uint16) (string, string) {
	info := pool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := pool.getUtf8(info.nameIndex)
	typE := pool.getUtf8(info.descriptorIndex)
	return name, typE
}
/* getClassName seek for class name from constant pool */
func (pool ConstantPool) getClassName(index uint16) string {
	info := pool.getConstantInfo(index).(*ConstantClassInfo)
	return pool.getUtf8(info.nameIndex)
}
/* getUtf8 seek for utf-8 string from constant pool */
func (pool *ConstantPool) getUtf8(index uint16) string {
	info := pool.getConstantInfo(index).(*ConstantUtf8Info)
	return info.str
}