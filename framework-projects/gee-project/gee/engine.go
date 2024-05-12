package gee

import (
	"net/http"
	"strings"
)

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

	groupHandles := make([]HandlerFunc, 0, len(e.groups))

	for _, group := range e.groups {
		if strings.HasPrefix(ctx.Path, group.prefix) {
			groupHandles = append(groupHandles, group.middlewares...)
		}
	}

	ctx.middlewares = groupHandles

	e.route.handle(ctx)
}
