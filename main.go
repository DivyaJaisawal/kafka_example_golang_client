package main

import (
	"fmt"
	"kafka_example_golang_client/consul_discover"
	server1 "kafka_example_golang_client/server"
	"log"
)

func main() {
	uServiceConsul, err := consul_discover.NewConsulClient("localhost:8500")
	if err != nil {
		log.Fatalln("Can't find consul_discover:", err)
	}
	r := server1.Router(uServiceConsul)
	server1.Start(r)
	fmt.Println("/greet endpoint is called")
}