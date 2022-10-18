package elysium

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"
)

//go:generate ./scripts/generate_version.sh
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

// new server instance
func New() *Engine {
	// init router
	var engine *Engine = &Engine{
		router: map[string]HandlerFunc{},
	}
	return engine
}

// server handle entry
func (engine *Engine) ServeHTTP(hw http.ResponseWriter, hr *http.Request) {
	log.Printf("%s - - [%s] \"%s %s %s\" - -", hr.RemoteAddr, time.Now().Format("2006/01/02 15:04:05"), hr.Method, hr.URL.Path, hr.Proto)
	key := strings.Join([]string{hr.Method, hr.URL.Path}, "-")
	if handler, ok := engine.router[key]; ok {
		handler(hw, hr)
	} else {
		fmt.Fprintf(hw, "404 NOT FOUND: %s\n", hr.URL)
	}
}

// server handle router: add
func (engine *Engine) addRoute(method, pattern string, handlers HandlerFunc) {
	key := method + "-" + pattern
	handlerName := runtime.FuncForPC(reflect.ValueOf(handlers).Pointer()).Name()
	log.Printf("%-4s %-25s ---> %s", method, pattern, handlerName)
	engine.router[key] = handlers
}
