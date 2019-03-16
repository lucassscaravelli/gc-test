package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RunServer cria uma nova instancia do router e starta o servidor
func RunServer() {
	mux := mux.NewRouter()

	server := newServer(mux)
	http.ListenAndServe(":8080", server)
}
