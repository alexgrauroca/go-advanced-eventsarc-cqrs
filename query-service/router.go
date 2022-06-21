package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/feeds", listFeedsHandler).Methods(http.MethodGet)
	router.HandleFunc("/search", searchHandler).Methods(http.MethodGet)

	return
}
