package classfile

import (
	"fmt"
	"unicode/utf16"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/13 20:56
 */

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	str string
}

func (info *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	info.str = decodeMUTF8(bytes)
}

func (info *ConstantUtf8Info) Str() string {
	return info.str
}

/*
func decodeMUTF8(bytes []byte) string {
	return string(bytes) // not correct!
}
*/

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(byteArr []byte) string {
	utfLen := len(byteArr)
	charArr := make([]uint16, utfLen)

	var c, char2, char3 uint16
	count := 0
	charArrCount := 0

	for count < utfLen {
		c = uint16(byteArr[count])
		if c > 127 {
			break
		}
		count++
		charArr[charArrCount] = c
		charArrCount++
	}

	for count < utfLen {
		c = uint16(byteArr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			charArr[charArrCount] = c
			charArrCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(byteArr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			charArr[charArrCount] = c&0x1F<<6 | char2&0x3F
			charArrCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(byteArr[count-2])
			char3 = uint16(byteArr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count - 1))
			}
			charArr[charArrCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charArrCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utfLen
	charArr = charArr[0:charArrCount]
	runes := utf16.Decode(charArr)
	return string(runes)
}
