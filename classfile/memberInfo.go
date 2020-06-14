package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 15:51
 */

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type MemberInfo struct {
	pool            ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}
/* readMembers read the table of fields and methods */
func readMembers(reader *ClassReader, pool ConstantPool) []*MemberInfo {
	count := reader.readUint16()
	members := make([]*MemberInfo, count)
	for i := range members {
		members[i] = readMember(reader, pool)
	}
	return members
}
/* readMember read the data of field and method */
func readMember(reader *ClassReader, pool ConstantPool) *MemberInfo {
	return &MemberInfo{
		pool:            pool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, pool),
	}
}
/* getter for access flags */
func (info *MemberInfo) AccessFlags() uint16 {
	return info.accessFlags
}
/* Name seek for the name of field or method from constant pool */
func (info *MemberInfo) Name() string {
	return info.pool.getUtf8(info.nameIndex)
}
/* Descriptor seek for the descriptor of field or method from constant pool */
func (info *MemberInfo) Descriptor() string {
	return info.pool.getUtf8(info.descriptorIndex)
}