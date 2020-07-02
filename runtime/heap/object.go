package heap

/**
 * @Author: Zhou Zilong
 * @Date: 2020/6/14 16:27
 */

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (object *Object) Class() *Class {
	return object.class
}
func (object *Object) Fields() Slots {
	return object.fields
}

func (object *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(object.class)
}
