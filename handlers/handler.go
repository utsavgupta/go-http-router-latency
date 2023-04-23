package handlers

import "net/http"

func NewHandler(path string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "text/plain;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(path))
	}
}
