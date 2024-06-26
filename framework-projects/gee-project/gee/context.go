package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]any

type HandlerFunc func(*Context)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string            // URL: /hello/maodou
	Method string            // GET/POST/DELETE等
	Params map[string]string // URL: /hello/:name  --> { name: value }

	// response info
	StatusCode int

	middlewares []HandlerFunc
	idx         int // init -1, no handler func exec
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,

		idx: -1,
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)

	_, _ = c.Writer.Write([]byte(html))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) String(code int, str string, format any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, _ = c.Writer.Write([]byte(fmt.Sprintf(str, format)))
}

func (c *Context) Data(code int, b []byte) {
	c.Status(code)
	_, _ = c.Writer.Write(b)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(c.StatusCode)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) Next() {
	c.idx++

	for ; c.idx < len(c.middlewares); c.idx++ {
		c.middlewares[c.idx](c)
	}
}
