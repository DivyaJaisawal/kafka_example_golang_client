package mock_greet

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	pb "kafka_example_golang_client/proto"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Greeting(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	args := m.Called(ctx, req)
	return args[0].(*pb.HelloResponse), args.Error(1)
}
