package handler_test

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"kafka_example_golang_client/mocks"
	mockConsulClient "kafka_example_golang_client/mocks"
	pb "kafka_example_golang_client/proto"
	http_server "kafka_example_golang_client/server"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func startService() (*grpc.Server, net.Listener, *mocks.MockService) {
	server := grpc.NewServer()
	svc := &mocks.MockService{}
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

	mockService.On("Greeting", mock.Anything, mock.Anything).Return(&pb.HelloResponse{
		Greeting: "hello",
	}, nil)

	mockConsulClient := &mockConsulClient.MockConsulClient{}
	mockConsulClient.On("Service", "kafka-example",
		"kafka").Return([]*consul.ServiceEntry{{
		Node: &consul.Node{
			ID:              "localhost",
			Node:            "50051",
			Address:         "localhost",
			Datacenter:      "dc1",
			TaggedAddresses: nil,
			Meta:            nil,
			CreateIndex:     0,
			ModifyIndex:     0,},
		Service: &consul.AgentService{Port: 50051},
	}}, &consul.QueryMeta{}, nil)

	router := http_server.Router(mockConsulClient)
	server := httptest.NewServer(router)
	defer server.Close()

	request, err := http.NewRequest("POST", server.URL+"/greet", strings.NewReader(`{"message": "Hello Divya Jaisawal"}`))

	if err != nil {
		fmt.Printf("error in connecting to http request: %v", err)
	}
	client := http.Client{}

	response, err := client.Do(request)
	fmt.Printf("got success response : %v", response)
	assert.NotNil(t, response)
	assert.Equal(t, 200,response.StatusCode)
}
