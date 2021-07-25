package event

import "reflect"

type Registry struct {
	handlerRegistry map[eventType][]handlerFunction
}

func NewRegistry() *Registry {
	return &Registry{
		handlerRegistry: make(map[eventType][]handlerFunction),
	}
}

func (r *Registry) RegistrerHandlerFunc(handlerFuncObject interface{}) *Registry {
	handlerFunc := reflect.ValueOf(handlerFuncObject)
	handlerFuncType := handlerFunc.Type()

	if handlerFuncType.NumIn() != 1 {
		panic("handler function must have a single argument as an event")
	}

	typeOfArgumentAsEvent := handlerFuncType.In(0)

	r.addHandlerFunc(typeOfArgumentAsEvent, handlerFunc)

	return r
}

func (r *Registry) addHandlerFunc(t eventType, handlerFunc handlerFunction) {
	handlerFuncs, ok := r.handlerRegistry[t]
	if !ok {
		handlerFuncs = make([]handlerFunction, 0)
	}

	r.handlerRegistry[t] = append(handlerFuncs, handlerFunc)
}

func (r *Registry) Bus() Bus {
	return &bus{
		handlerRegistry: r.handlerRegistry,
	}
}
