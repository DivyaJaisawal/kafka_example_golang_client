package handler_test

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"kafka_example_golang_client/mock_greet"
	pb "kafka_example_golang_client/proto"
	http_server "kafka_example_golang_client/server"
	"strings"
	"testing"
)

func startService() (*grpc.Server, net.Listener, *mock_greet.MockService) {
	server := grpc.NewServer()
	svc := &mock_greet.MockService{}
	pb.RegisterGreetingServiceServer(server, svc)
	listener, _ := net.Listen("tcp", "localhost:50051")
	go func() {
		server.Serve(listener)
	}()
	return server, listener, svc
}

func TestShouldReturn200ForSuccessResponse(t *testing.T) {
	_, listener, mockService := startService()
	defer listener.Close()

	mockService.On("Greeting", mock.Anything, mock.Anything).Return(&pb.HelloResponse {
		Greeting: "hello",
	}, nil)

	router := http_server.Router()
	http_server.Start(router)

	request, err := http.NewRequest("POST", "http://localhost:8080/greet", strings.NewReader(`{"message": "Hello Divya Jaisawal"}`))

	if err != nil {
		fmt.Printf("error in connecting to http request: %v", err)
	}
	client := http.Client{}

	response, err := client.Do(request)
	fmt.Printf("got success response : %v", response)

	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetingServiceClient(conn)
	t.Run("Greeting", func(t *testing.T) {
		name := "hello"
		r, err := c.Greeting(context.Background(), &pb.HelloRequest{Message: name})
		if err != nil {
			t.Fatalf("could not greet: %v", err)
		}
		t.Logf("Greeting: %s", r.Greeting)
		if r.Greeting != "Hello "+name {
			t.Error("Expected 'Hello world', got ", r.Greeting)
		}
	})
}
