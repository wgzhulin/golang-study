package main

import (
	"fmt"
	"gostudy/framework-projects/gee-project/gee"
	"net/http"
)

func main() {
	g := gee.New()

	g.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	g.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello %v", ctx.Query("name"))
	})

	g.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	v2 := g.Group("/v2")
	v2.GET("/hello", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, fmt.Sprintf("<h1>Hello %s</h1>", ctx.Query("name")))
	})

	http.ListenAndServe(":8080", g)
}
