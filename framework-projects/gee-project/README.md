# WEB框架之实现简单的Gin框架

[七天用Go从零实现系列](https://geektutu.com/post/gee.html)学习

学习目标
1. go标准库如何启动Web服务
2. Router设计，路由处理，动态路由，分组控制
3. 中间件

net/http下接口`ServeHTTP`，Go标准库中`ServeMux`默认实现该接口。也可以自定义实现改接口，调用`ListenAndServe`传入自定义实现，接管所有的HTTP请求
