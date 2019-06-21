package server1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start(r *mux.Router)  {
	http.ListenAndServe("localhost:8080", r)
}