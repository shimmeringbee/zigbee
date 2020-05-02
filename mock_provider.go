package zigbee

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockProvider struct {
	mock.Mock
}

func (m *MockProvider) AdapterNode() Node {
	args := m.Called()
	return args.Get(0).(Node)
}

func (m *MockProvider) QueryNodeDescription(ctx context.Context, networkAddress IEEEAddress) (NodeDescription, error) {
	args := m.Called(ctx, networkAddress)
	return args.Get(0).(NodeDescription), args.Error(1)
}

func (m *MockProvider) QueryNodeEndpoints(ctx context.Context, networkAddress IEEEAddress) ([]byte, error) {
	args := m.Called(ctx, networkAddress)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockProvider) QueryNodeEndpointDescription(ctx context.Context, networkAddress IEEEAddress, endpoint Endpoint) (EndpointDescription, error) {
	args := m.Called(ctx, networkAddress, endpoint)
	return args.Get(0).(EndpointDescription), args.Error(1)
}

func (m *MockProvider) BindNodeToController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint Endpoint, destinationEndpoint Endpoint, cluster ClusterID) error {
	args := m.Called(ctx, nodeAddress, sourceEndpoint, destinationEndpoint, cluster)
	return args.Error(0)
}

func (m *MockProvider) UnbindNodeFromController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint Endpoint, destinationEndpoint Endpoint, cluster ClusterID) error {
	args := m.Called(ctx, nodeAddress, sourceEndpoint, destinationEndpoint, cluster)
	return args.Error(0)
}

func (m *MockProvider) SendNodeMessageToNode(ctx context.Context, destinationAddress IEEEAddress, message ApplicationMessage) error {
	args := m.Called(ctx, destinationAddress, message)
	return args.Error(0)
}

func (m *MockProvider) ReadEvent(ctx context.Context) (interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0), args.Error(1)
}