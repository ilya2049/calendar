package event

import "reflect"

type Bus interface {
	Publish(eventObject interface{})
}

type bus struct {
	handlerRegistry map[eventType][]handlerFunction
}

func (b *bus) Publish(eventObject interface{}) {
	t := reflect.TypeOf(eventObject)

	if handlerFuncs, ok := b.handlerRegistry[t]; ok {
		handlerFuncArguments := [...]handlerFunction{reflect.ValueOf(eventObject)}
		for _, handlerFunc := range handlerFuncs {
			handlerFunc.Call(handlerFuncArguments[:])
		}
	}
}
