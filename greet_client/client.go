package greet_client

import (
	"fmt"
	"google.golang.org/grpc"
	"kafka_example_golang_client/consul_discover"
	greet "kafka_example_golang_client/proto"
	"log"
	"strconv"
)

func GetClient() greet.GreetingServiceClient{
	uServiceConsul, err := consul_discover.NewConsulClient("localhost:8500")
	if err != nil {
		log.Fatalln("Can't find consul_discover:", err)
	}
	services, _, _ := uServiceConsul.Service("kafka-example", "kafka")
	if err != nil {
		log.Fatalln("Discover failed:", err)
	}
	var conn *grpc.ClientConn
	var grpcErr error
	log.Println("Found service at these locations:")
	for _, v := range services {
		log.Println(fmt.Sprintf("%s:%d", v.Node.Address, v.Service.Port))

		conn, grpcErr = grpc.Dial(v.Node.Address+":"+ strconv.Itoa(v.Service.Port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", grpcErr)
		}
	}
	return greet.NewGreetingServiceClient(conn)
}
