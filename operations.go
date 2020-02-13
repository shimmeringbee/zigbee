package zigbee

import (
	"context"
	"time"
)

type NodeQuerier interface {
	QueryNodeDescription(ctx context.Context, networkAddress IEEEAddress) (NodeDescription, error)
	QueryNodeEndpoints(ctx context.Context, networkAddress IEEEAddress) ([]byte, error)
	QueryNodeEndpointDescription(ctx context.Context, networkAddress IEEEAddress, endpoint Endpoint) (EndpointDescription, error)
}

type NodeBinder interface {
	BindNodeToController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint Endpoint, destinationEndpoint Endpoint, cluster ClusterID) error
	UnbindNodeFromController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint Endpoint, destinationEndpoint Endpoint, cluster ClusterID) error
}

type NodeSender interface {
	SendNodeMessage(ctx context.Context, destinationAddress IEEEAddress, sourceEndpoint Endpoint, destinationEndpoint Endpoint, cluster ClusterID, data []byte) error
}

type EventReceiver interface {
	ReadEvent(ctx context.Context) (interface{}, error)
}

type Node struct {
	IEEEAddress    IEEEAddress
	NetworkAddress NetworkAddress
	LogicalType    LogicalType
	LQI            uint8
	Depth          uint8
	LastDiscovered time.Time
	LastReceived   time.Time
}

type NodeJoinEvent struct {
	Node
}

type NodeUpdateEvent struct {
	Node
}

type NodeLeaveEvent struct {
	Node
}

type IncomingMessage struct {
	GroupID              GroupID
	ClusterID            ClusterID
	SourceIEEEAddress    IEEEAddress
	SourceNetworkAddress NetworkAddress
	SourceEndpoint       Endpoint
	DestinationEndpoint  Endpoint
	Broadcast            bool
	Secure               bool
	LinkQuality          uint8
	Sequence             uint8
	Data                 []byte
}

type NodeIncomingMessageEvent struct {
	Node
	IncomingMessage
}
