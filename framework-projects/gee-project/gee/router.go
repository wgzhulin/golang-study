package gee

import "net/http"

type Router struct {
	m map[string]HandlerFunc
}

func newRouter() Router {
	return Router{m: make(map[string]HandlerFunc, 0)}
}

func (r *Router) addRoute(method string, pattern string, function HandlerFunc) {
	key := method + "-" + pattern
	r.m[key] = function
}

func (r *Router) handle(pattern string, ctx *Context) {
	if handler, ok := r.m[pattern]; ok {
		handler(ctx)
	} else {
		http.NotFound(ctx.Writer, ctx.Req)
	}
}
