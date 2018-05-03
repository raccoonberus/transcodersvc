package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func TaskCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	w.Write([]byte(filename)) // TODO: implement method
}


func TaskInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	w.Write([]byte(filename)) // TODO: implement method
}
