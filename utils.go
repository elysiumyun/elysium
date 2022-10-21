package elysium

import (
	"reflect"
	"runtime"
)

func reflectHandlerName(handler HandlerFunc) string {
	handlerName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return handlerName
}
