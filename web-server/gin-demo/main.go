package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// GET /path/sum?a=1&b=2
	engine.GET("/sum", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")
		intA, _ := strconv.Atoi(a)
		intB, _ := strconv.Atoi(b)

		c.JSON(http.StatusOK, fmt.Sprintf("sum: %v", intA+intB))
	})

	// GET /path/hello/wzl
	engine.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name") // name = "wzl"
		c.JSON(http.StatusOK, fmt.Sprintf("hello %v", name))
	})

	// POST /path/login
	// form data: {
	// 		user   	   wzl
	// 		password   wzlpass
	// }
	engine.POST("/login", func(c *gin.Context) {
		user := c.PostForm("user")         // user = "wzl"
		password := c.PostForm("password") // password = "wzlpass"

		var result string
		if user == "wzl" && password == "wzlpass" {
			result = "login success"
		} else {
			result = "login fail, user or password error"
		}
		c.JSON(http.StatusOK, result)
	})

	engine.Run(":80")
}
