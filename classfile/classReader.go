package classfile

import "encoding/binary"

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/12 9:51
 */

type ClassReader struct {
	data []byte
}

/* different read methods of class reader */

// u1
func (reader *ClassReader) readUint8() uint8 {
	val := reader.data[0]
	reader.data = reader.data[1:]
	return val
}
// for u2
func (reader *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(reader.data)
	reader.data = reader.data[2:]
	return val
}
// for u4
func (reader *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(reader.data)
	reader.data = reader.data[4:]
	return val
}
// for u8
func (reader *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(reader.data)
	reader.data = reader.data[8:]
	return val
}
// for u2 table
func (reader *ClassReader) readUint16s() []uint16 {
	n := reader.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = reader.readUint16()
	}
	return s
}
// for byte
func (reader *ClassReader) readBytes(n uint32) []byte {
	bytes := reader.data[:n]
	reader.data = reader.data[n:]
	return bytes
}