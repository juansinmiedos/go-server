package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return r
}