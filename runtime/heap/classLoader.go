package heap

import (
	"fmt"
	"komorebi/classfile"
	"komorebi/classpath"
)

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/16 20:25
 */

/**
 * you see that the load of classes could be devided into three steps:
 *   first we find the class files and load the data to the memory
 *   then we parse the class files and generate the data that could be used by JVM and put it to map
 *   and finally we do the link
 */

type ClassLoader struct {
	cp          *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
}

func (loader *ClassLoader) LoadClass(name string) *Class {
	if class, ok := loader.classMap[name]; ok {
		// already loaded
		return class
	}

	return loader.loadNonArrayClass(name)
}

/* array class is different from normal class, its data doesn't come from class file, but from the
running JVM, so here we don't think about loading array class */
func (loader *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := loader.readClass(name)
	class := loader.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

/* readClass simply call the ReadClass of classpath, and to print the class message, the entry is
also returned to the caller */
func (loader *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := loader.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

/* defineClass first call the parseClass to turn the class file into the Class struct, in which
there are fields like superClass and interfaces to store the table of super class and interface,
which are all the references */
func (loader *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = loader
	resolveSuperClass(class)
	resolveInterfaces(class)
	loader.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

/* here we don't focus on verifying */
func verify(class *Class) {
	// nothing to do
}

// jvms 5.4.2
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
