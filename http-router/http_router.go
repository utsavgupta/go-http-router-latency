package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/utsavgupta/go-http-router-latency/handlers"
)

func main() {

	fmt.Println("Starting server with julienschmidt httprouter")

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", handlers.Index)
	router.HandlerFunc(http.MethodGet, "/about", handlers.About)
	router.HandlerFunc(http.MethodGet, "/projects", handlers.Projects)
	router.HandlerFunc(http.MethodGet, "/posts", handlers.PostsIndex)
	router.HandlerFunc(http.MethodGet, "/posts/:pid", handlers.PostView)
	router.HandlerFunc(http.MethodGet, "/posts/:pid/comments", handlers.CommentsIndex)
	router.HandlerFunc(http.MethodGet, "/posts/:pid/comments/:cid", handlers.CommentView)

	http.ListenAndServe(":8080", router)
}
