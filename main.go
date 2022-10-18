package elysium

import (
	"net/http"
)

//go:generate ./scripts/generate_version.sh
type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

// new server instance
func New() *Engine {
	// init router
	var engine *Engine = &Engine{
		router: newRouter(),
	}
	return engine
}

// server handle router: add
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// server handle entry
func (engine *Engine) ServeHTTP(hw http.ResponseWriter, hr *http.Request) {
	ctx := newContext(hw, hr)
	engine.router.handle(ctx)
}
