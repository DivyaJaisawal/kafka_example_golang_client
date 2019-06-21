package handler_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	greet "source.golabs.io/gopay_apps/kafka_example_golang_client/proto"
	server1 "source.golabs.io/gopay_apps/kafka_example_golang_client/server"
	"strings"
	"testing"
)

func TestShouldReturn200ForSuccessResponse(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := greet.NewGreetingServiceClient(conn)

	// Test SayHello
	t.Run("SayHello", func(t *testing.T) {
		name := "world"
		r, err := c.Greeting(context.Background(), &greet.HelloRequest{Message: name})
		if err != nil {
			t.Fatalf("could not greet: %v", err)
		}
		t.Logf("Greeting: %s", r.Greeting)
		if r.Greeting != "Hello "+name {
			t.Error("Expected 'Hello world', got ", r.Greeting)
		}

	})
}

type MockGreetService struct {
	mock.Mock
}

func (m *MockGreetService) Greeting(ctx context.Context, r *greet.HelloRequest) (*greet.HelloResponse, error) {
	args := m.Called(ctx, r)
	return args[0].(*greet.HelloResponse), args.Error(1)
}

type TestContext struct {
	t          *testing.T
	deferFuncs []func()
}

func NewTestContext(t *testing.T) *TestContext {
	return &TestContext{
		t: t,
	}
}

func TestApi(t *testing.T) {

	//start the http server
	router := server1.Router()
	server1.Start(router)
	//make a http
	request := httptest.NewRequest("POST", "/localhost:8080/greet", strings.NewReader(`{"message": "Hello Divya Jaisawal"}`))
	// create a client
	client := http.Client{}
	// make the request
	response, _ := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)

/*	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := greet.NewGreetingServiceClient(conn)

	// Test SayHello
	t.Run("SayHello", func(t *testing.T) {
		name := "world"
		r, err := c.Greeting(context.Background(), &greet.HelloRequest{Message: name})
		if err != nil {
			t.Fatalf("could not greet: %v", err)
		}
		t.Logf("Greeting: %s", r.Greeting)
		if r.Greeting != "Hello "+name {
			t.Error("Expected 'Hello world', got ", r.Greeting)
		}

	})
*/
/*
	client := greet.NewGreetingServiceClient(server.URL, hystrix.NewClient(), &logrus.Logger{})

	request := CancelReservationRequest{
		ReservationIdentifier: "rid",
		WalletID:              "wid",
		MerchantId:            "someMerchantId",
		RequestID:             "request_id",
	}

	response, err := client.CancelReservation(request, nil)

	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.Empty(t, response.Errors)
	assert.JSONEq(t, `{"status": true}`, string(bytes))*/
}
