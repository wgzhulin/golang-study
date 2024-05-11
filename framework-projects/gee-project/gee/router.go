package gee

import (
	"net/http"
	"strings"
)

// Router 接管HTTP请求，前缀树结构
type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() Router {
	return Router{
		roots:    make(map[string]*node, 0),
		handlers: make(map[string]HandlerFunc, 0),
	}
}

func (r *Router) addRoute(method string, pattern string, function HandlerFunc) {
	key := method + "-" + pattern

	parts := parsePattern(pattern)
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = function
}

func (r *Router) getRoute(method string, pattern string) (*node, map[string]string) {
	searchParts := parsePattern(pattern)

	params := make(map[string]string, 0)
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

func (r *Router) handle(ctx *Context) {
	n, params := r.getRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Req.Method + "-" + ctx.Req.URL.Path
		r.handlers[key](ctx)
	} else {
		http.NotFound(ctx.Writer, ctx.Req)
	}
}
