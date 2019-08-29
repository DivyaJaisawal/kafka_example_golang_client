package mocks

import "github.com/stretchr/testify/mock"
import consul "github.com/hashicorp/consul/api"
type MockConsulClient struct {
	mock.Mock
}

func (m *MockConsulClient) Service(a string, b string)  ([]*consul.ServiceEntry, *consul.QueryMeta, error){
	args := m.Called(a, b)
	return args[0].([]*consul.ServiceEntry), args[1].(*consul.QueryMeta), args.Error(2)
}

func (m *MockConsulClient) DeRegister(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockConsulClient) Register(name string,value int) error {
	args := m.Called(name)
	return args.Error(0)
}
