package greet_client

import (
	"fmt"
	"google.golang.org/grpc"
	"kafka_example_golang_client/consul_discover"
	greet "kafka_example_golang_client/proto"
	"log"
	"strconv"
)

func GetClient(uServiceConsul consul_discover.Client) greet.GreetingServiceClient{
	services, _, err := uServiceConsul.Service("kafka-example", "kafka")
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
	if(conn == nil) {
		panic("connection to grpc server failed")
	}
	return greet.NewGreetingServiceClient(conn)
}
