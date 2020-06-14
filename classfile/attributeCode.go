package classfile

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 10:49
 */

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	pool           ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (attr *CodeAttribute) readInfo(reader *ClassReader) {
	attr.maxStack = reader.readUint16()
	attr.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	attr.code = reader.readBytes(codeLength)
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.pool)
}

func (attr *CodeAttribute) MaxStack() uint {
	return uint(attr.maxStack)
}
func (attr *CodeAttribute) MaxLocals() uint {
	return uint(attr.maxLocals)
}
func (attr *CodeAttribute) Code() []byte {
	return attr.code
}
func (attr *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return attr.exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (entry *ExceptionTableEntry) StartPc() uint16 {
	return entry.startPc
}
func (entry *ExceptionTableEntry) EndPc() uint16 {
	return entry.endPc
}
func (entry *ExceptionTableEntry) HandlerPc() uint16 {
	return entry.handlerPc
}
func (entry *ExceptionTableEntry) CatchType() uint16 {
	return entry.catchType
}
