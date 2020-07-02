package heap

import (
	"komorebi/classfile"
	"strings"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 20:08
 */

type Class struct {
	accessFlags       uint16
	name              string // thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) IsPublic() bool {
	return 0 != class.accessFlags&ACC_PUBLIC
}
func (class *Class) IsFinal() bool {
	return 0 != class.accessFlags&ACC_FINAL
}
func (class *Class) IsSuper() bool {
	return 0 != class.accessFlags&ACC_SUPER
}
func (class *Class) IsInterface() bool {
	return 0 != class.accessFlags&ACC_INTERFACE
}
func (class *Class) IsAbstract() bool {
	return 0 != class.accessFlags&ACC_ABSTRACT
}
func (class *Class) IsSynthetic() bool {
	return 0 != class.accessFlags&ACC_SYNTHETIC
}
func (class *Class) IsAnnotation() bool {
	return 0 != class.accessFlags&ACC_ANNOTATION
}
func (class *Class) IsEnum() bool {
	return 0 != class.accessFlags&ACC_ENUM
}

// getters
func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) StaticVars() Slots {
	return class.staticVars
}

// jvms 5.4.4
func (class *Class) isAccessibleTo(other *Class) bool {
	return class.IsPublic() ||
		class.getPackageName() == other.getPackageName()
}

func (class *Class) getPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) GetMainMethod() *Method {
	return class.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (class *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}

func (class *Class) Name() string {
	return class.name
}
