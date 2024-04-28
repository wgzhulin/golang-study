package engine

import (
	"fmt"
	"net/http"
)

type Engine struct {
	routes map[string]http.HandlerFunc
}

func New() *Engine {
	return &Engine{routes: make(map[string]http.HandlerFunc, 0)}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if handler, ok := e.routes[req.Method+"-"+req.URL.Path]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL.Path)
	}
}

func (e *Engine) addRoute(method string, pattern string, handler http.HandlerFunc) {
	key := method + "-" + pattern
	e.routes[key] = handler
}

func (e *Engine) Get(pattern string, handler http.HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler http.HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
