package elysium

import (
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	handlerName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	log.Printf("%-4s %-25s ---> %s", method, pattern, handlerName)
	r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
	log.Printf("%s - - [%s] \"%s %s %s\" - -",
		ctx.R.RemoteAddr,
		time.Now().Format(time.ANSIC),
		ctx.R.Method,
		ctx.R.URL.Path,
		ctx.R.Proto,
	)
	key := strings.Join([]string{ctx.R.Method, ctx.R.URL.Path}, "-")
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 not found :%s\n", ctx.R.URL)
	}
}
