package main

import (
	"gostudy/7days-projects/project1-web-frame/day3-router/gee"
	"net/http"
)

func main() {
	g := gee.New()
	
	g.Get("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})


	g.Get("/hello/:name", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello %s", ctx.Params("name"))
	})

	g.Get("/assets/*filepath", func(ctx *gee.Context) {
		ctx.Json(http.StatusOK, gee.H{"filepath": ctx.Params("filepath")})
	})

	g.Run(":8080")
}