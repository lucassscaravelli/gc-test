package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer() {
	mux := mux.NewRouter()

	server := newServer(mux)
	http.ListenAndServe(":8080", server)
}
