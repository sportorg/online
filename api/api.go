package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run(addr string, pathToPWA string) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(pathToPWA)))
	log.Fatal(http.ListenAndServe(addr, r))
}
