package elysium

import (
	"log"
	"net/http"
	"strings"
	"time"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	handlerName := reflectHandlerName(handler)
	log.Printf("%-4s %-25s ---> %s", method, pattern, handlerName)

	parts := parsePattern(pattern)

	key := strings.Join([]string{method, pattern}, "-")
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *router) handle(ctx *Context) {
	log.Printf("%s - - [%s] \"%s %s %s\" - -",
		ctx.R.RemoteAddr,
		time.Now().Format(time.ANSIC),
		ctx.R.Method,
		ctx.R.URL.Path,
		ctx.R.Proto,
	)

	node, params := r.getRoute(ctx.R.Method, ctx.R.URL.Path)

	if node != nil {
		ctx.Params = params
		key := strings.Join([]string{ctx.R.Method, node.pattern}, "-")
		r.handlers[key](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 not found :%s\n", ctx.Path)
	}
}
