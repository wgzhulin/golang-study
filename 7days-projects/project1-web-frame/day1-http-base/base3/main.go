package main

import (
	"fmt"
	"gostudy/7days-projects/project1-web-frame/day1-http-base/base3/engine"
	"net/http"
)

func main() {
	e := engine.New()

	e.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	e.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Head[%q] = %q\n", k, v)
		}
	})

	e.Run(":8089")
}