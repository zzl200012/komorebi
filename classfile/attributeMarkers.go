package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 10:51
 */

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (attr *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
