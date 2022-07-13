package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	r := mux.NewRouter

	return r
}