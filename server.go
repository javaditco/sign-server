package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/upload", UploadHandler).Methods("PUT")
	r.HandleFunc("/file/{id}", GetFileHandler).Methods("GET")
	n.UseHandler(r)
	n.Run("0.0.0.0:8080")
}
