package main

import (
	"gostudy/7days-projects/project1-web-frame/day2-context/gee"
	"net/http"
)

func main() {
	h := gee.New()

	h.Get("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	h.Get("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello %s", ctx.Query("name"))
	})

	h.POST("/login", func(ctx *gee.Context) {
		ctx.Json(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	h.Run(":8082")
}
