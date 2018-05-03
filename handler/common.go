package handler

import "net/http"

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

type boostedHandlerFunc func(w http.ResponseWriter, r *http.Request, argv ... interface{})

type BoostedMiddleWare func(handlerFunc boostedHandlerFunc) http.HandlerFunc
