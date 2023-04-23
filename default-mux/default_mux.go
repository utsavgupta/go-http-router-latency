package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/utsavgupta/go-http-router-latency/handlers"
)

func main() {

	fmt.Println("Starting server with default router")

	mux := http.DefaultServeMux

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/projects", handlers.Projects)
	mux.HandleFunc("/posts/", postsHandler(handlers.PostsIndex, handlers.PostView, handlers.CommentsIndex, handlers.CommentView))

	http.ListenAndServe(":8080", mux)
}

func postsHandler(index, view, commentsIndex, commentsView http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		re, _ := regexp.Compile(`/posts(/(\d+)(/comments(/(\d+))?)?)?`)

		matches := re.FindStringSubmatch(path)

		_, commentIdParseErr := strconv.Atoi(matches[5])
		_, postIdParseErr := strconv.Atoi(matches[2])

		if commentIdParseErr == nil {
			commentsView(w, r)
			return
		}

		if len(matches[3]) > 0 {
			commentsIndex(w, r)
			return
		}

		if postIdParseErr == nil {
			view(w, r)
			return
		}

		index(w, r)
	}
}
