package server1

import (
	"github.com/gorilla/mux"
	"kafka_example_golang_client/consul_discover"
	"kafka_example_golang_client/greet_client"
	"kafka_example_golang_client/handler"
)

func Router(consulcli consul_discover.Client) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/greet", handler.GreetHandler(greet_client.GetClient(consulcli))).Methods("POST")
	return r
}
