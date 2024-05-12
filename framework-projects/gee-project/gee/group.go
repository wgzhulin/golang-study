package gee

import "log"

type RouterGroup struct {
	prefix      string
	parent      *RouterGroup
	engine      *Engine
	middlewares []HandlerFunc
}

func (r *RouterGroup) Use(middlewares ...HandlerFunc) {
	r.middlewares = middlewares
}

func (r *RouterGroup) Group(prefix string) *RouterGroup {
	e := r.engine
	newGroup := &RouterGroup{
		prefix: prefix,
		parent: r,
		engine: e,
	}

	e.groups = append(e.groups, newGroup)
	return newGroup
}

func (r *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := r.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	r.engine.route.addRoute(method, pattern, handler)
}

func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}
