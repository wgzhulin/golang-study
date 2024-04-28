package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", new(Engine))
}


type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Head[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "invalid URL.Path = %q\n", req.URL.Path)
	}
}
