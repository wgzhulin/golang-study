package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("Index"))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("Hello World\n"))

		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Head[%q] = %q\n", k, v)
		}
	})

	http.HandleFunc("/api/:name", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("match name\n"))
	})

	_ = http.ListenAndServe(":8080", nil)
}
