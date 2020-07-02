package heap

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/17 11:29
 */

type Ref struct {
	pool      *ConstantPool
	className string
	class     *Class
}

func (ref *Ref) ResolvedClass() *Class {
	if ref.class == nil {
		ref.resolveClassRef()
	}
	return ref.class
}

// jvms8 5.4.3.1
func (ref *Ref) resolveClassRef() {
	d := ref.pool.class
	c := d.loader.LoadClass(ref.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	ref.class = c
}
