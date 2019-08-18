package main

import (
	"fmt"
	server1 "kafka_example_golang_client/server"
)

func main() {

	r := server1.Router()
	server1.Start(r)
	fmt.Println("/greet endpoint is called")
}