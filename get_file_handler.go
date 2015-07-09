package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetFileHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	file_id := vars["id"]
	http.ServeFile(res, req, fmt.Sprintf("signatures/%s", file_id))
}
