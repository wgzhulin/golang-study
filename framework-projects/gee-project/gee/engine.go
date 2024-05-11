package gee

import "net/http"

type Engine struct {
	RouterGroup
	route  Router
	groups []*RouterGroup
}

func New() *Engine {
	e := &Engine{route: newRouter()}
	e.engine = e
	return e
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	ctx.Method = req.Method
	ctx.Path = req.URL.Path

	e.route.handle(ctx)
}
