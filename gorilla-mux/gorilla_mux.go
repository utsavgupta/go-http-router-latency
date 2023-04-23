package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/utsavgupta/go-http-router-latency/handlers"
)

func main() {

	fmt.Println("Starting server with gorilla mux")

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/about", handlers.About)
	router.HandleFunc("/projects", handlers.Projects)
	router.HandleFunc("/posts", handlers.PostsIndex)
	router.HandleFunc("/posts/{pid}", handlers.PostView)
	router.HandleFunc("/posts/{pid}/comments", handlers.CommentsIndex)
	router.HandleFunc("/posts/{pid}/comments/{cid}", handlers.CommentView)

	http.ListenAndServe(":8080", router)
}
