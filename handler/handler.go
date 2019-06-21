package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"source.golabs.io/gopay_apps/kafka_example_golang_client/consul_discover"
	greet "source.golabs.io/gopay_apps/kafka_example_golang_client/proto"
	"strconv"
	"time"
)

func GreetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var greetRequest greet.HelloRequest
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&greetRequest);
		consulDiscovery()
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := greet.NewGreetingServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		response, err := client.Greeting(ctx, &greet.HelloRequest{Message: greetRequest.Message})

		log.Printf("Greeting: %s", response.Greeting)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":" + strconv.FormatBool(response.Success)+ "}"))
	})
}

func consulDiscovery() error {
	uServiceConsul, err := consul_discover.NewConsulClient("localhost:8500")
	if err != nil {
		log.Fatalln("Can't find consul_discover:", err)
	}
	services, _, err := uServiceConsul.Service("kafka-example", "kafka")
	if err != nil {
		log.Fatalln("Discover failed:", err)
	}
	log.Println("Found service at these locations:")
	for _, v := range services {
		log.Println(fmt.Sprintf("%s:%d", v.Node.Address, v.Service.Port))
	}
	return err
}
