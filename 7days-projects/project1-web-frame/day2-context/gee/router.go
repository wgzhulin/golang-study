package gee

import "net/http"

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: map[string]HandlerFunc{}}
}

func (e *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.handlers[key] = handler
}

func (e *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if f, ok := e.handlers[key]; ok {
		f(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
