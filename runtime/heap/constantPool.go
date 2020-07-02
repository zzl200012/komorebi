package heap

import (
	"fmt"
	"komorebi/classfile"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 11:08
 */

type Constant interface{}

type ConstantPool struct {
	class *Class
	pool  []Constant
}

/* turn the constant pool of class file into runtime constant pool, in which we just turn the constant info
of class file into Constant struct */
func newConstantPool(class *Class, pool classfile.ConstantPool) *ConstantPool {
	length := len(pool)
	p := make([]Constant, length)
	cp := &ConstantPool{}
	cp.class = class
	cp.pool = p
	for i := 0; i < length; i++ {
		info := pool[i]
		switch info.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := info.(*classfile.ConstantIntegerInfo)
			p[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := info.(*classfile.ConstantFloatInfo)
			p[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := info.(*classfile.ConstantLongInfo)
			p[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := info.(*classfile.ConstantDoubleInfo)
			p[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := info.(*classfile.ConstantStringInfo)
			p[i] = stringInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo := info.(*classfile.ConstantClassInfo)
			p[i] = newClassRef(cp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldRefInfo := info.(*classfile.ConstantFieldRefInfo)
			p[i] = newFieldRef(cp, fieldRefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodRefInfo := info.(*classfile.ConstantMethodRefInfo)
			p[i] = newMethodRef(cp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			methodRefInfo := info.(*classfile.ConstantInterfaceMethodRefInfo)
			p[i] = newInterfaceMethodRef(cp, methodRefInfo)
		default:
			// todo
		}
	}
	return cp
}

/* seek for constant by index */
func (pool *ConstantPool) GetConstant(index uint) Constant {
	if c := pool.pool[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constant at index %d", index))
}
