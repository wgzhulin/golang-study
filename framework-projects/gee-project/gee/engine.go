package gee

import "net/http"

type Engine struct {
	route Router
}

func New() *Engine {
	return &Engine{route: newRouter()}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	ctx.Method = req.Method
	ctx.Path = req.URL.Path

	e.route.handle(ctx)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.route.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.route.addRoute("POST", pattern, handler)
}
