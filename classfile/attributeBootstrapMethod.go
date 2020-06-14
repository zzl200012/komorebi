package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 10:48
 */

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/
type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (attr *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	attr.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range attr.bootstrapMethods {
		attr.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}
