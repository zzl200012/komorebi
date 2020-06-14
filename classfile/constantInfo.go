package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 16:32
 */

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/

/* definitions for different constant types */
const (
	ConstantClass              = 7
	ConstantFieldRef           = 9
	ConstantMethodRef          = 10
	ConstantInterfaceMethodRef = 11
	ConstantString             = 8
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantNameAndType        = 12
	ConstantUtf8               = 1
	ConstantMethodHandle       = 15
	ConstantMethodType         = 16
	ConstantInvokeDynamic      = 18
)

/* readInfo read the info of constant, which should be implemented by specific constant struct.
readConstantInfo read the tag first, then call the newConstant to create specific constant, and
finally call the readInfo of the specific constant to read the info of it */
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}
func readConstantInfo(reader *ClassReader, pool ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, pool)
	c.readInfo(reader)
	return c
}
func newConstantInfo(tag uint8, pool ConstantPool) ConstantInfo {
	switch tag {
	case ConstantInteger:
		return &ConstantIntegerInfo{}
	case ConstantFloat:
		return &ConstantFloatInfo{}
	case ConstantLong:
		return &ConstantLongInfo{}
	case ConstantDouble:
		return &ConstantDoubleInfo{}
	case ConstantUtf8:
		return &ConstantUtf8Info{}
	case ConstantString:
		return &ConstantStringInfo{pool: pool}
	case ConstantClass:
		return &ConstantClassInfo{pool: pool}
	case ConstantFieldRef:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{pool: pool}}
	case ConstantMethodRef:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{pool: pool}}
	case ConstantInterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{pool: pool}}
	case ConstantNameAndType:
		return &ConstantNameAndTypeInfo{}
	case ConstantMethodType:
		return &ConstantMethodTypeInfo{}
	case ConstantMethodHandle:
		return &ConstantMethodHandleInfo{}
	case ConstantInvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: Invalid constant pool tag!")
	}
}