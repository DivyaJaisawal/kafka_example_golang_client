package server1

import (
	"github.com/gorilla/mux"
	"source.golabs.io/gopay_apps/kafka_example_golang_client/handler"
)

func Router() *mux.Router{
	r := mux.NewRouter()
	r.Handle("/greet",handler.GreetHandler()).Methods("POST")
	return r
}
