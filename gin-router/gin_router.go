package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/utsavgupta/go-http-router-latency/handlers"
)

func main() {

	fmt.Println("Starting server with gin")

	engine := gin.New()

	engine.GET("/", gin.WrapF(handlers.Index))
	engine.GET("/about", gin.WrapF(handlers.About))
	engine.GET("/projects", gin.WrapF(handlers.Projects))
	engine.GET("/posts", gin.WrapF(handlers.PostsIndex))
	engine.GET("/posts/:pid", gin.WrapF(handlers.PostView))
	engine.GET("/posts/:pid/comments", gin.WrapF(handlers.CommentsIndex))
	engine.GET("/posts/:pid/comments/:cid", gin.WrapF(handlers.CommentView))

	engine.Run(":8080")
}
