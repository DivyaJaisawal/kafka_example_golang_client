package main

import (
	"fmt"
	server1 "source.golabs.io/gopay_apps/kafka_example_golang_client/server"
)

func main() {

	r := server1.Router()
	server1.Start(r)
	fmt.Println("/greet endpoint is called")
}