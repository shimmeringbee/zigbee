package zigbee

import (
	"context"
	"time"
)

type NetworkJoining interface {
	PermitJoin(ctx context.Context, allRouters bool) error
	DenyJoin(ctx context.Context) error
}

type AdapterInfo interface {
	AdapterNode() Node
}

type NodeQuerier interface {
	QueryNodeDescription(ctx context.Context, networkAddress IEEEAddress) (NodeDescription, error)
	QueryNodeEndpoints(ctx context.Context, networkAddress IEEEAddress) ([]Endpoint, error)
	QueryNodeEndpointDescription(ctx context.Context, networkAddress IEEEAddress, endpoint Endpoint) (EndpointDescription, error)
}

type NodeBinder interface {
	BindNodeToController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint Endpoint, destinationEndpoint Endpoint, cluster ClusterID) error
	UnbindNodeFromController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint Endpoint, destinationEndpoint Endpoint, cluster ClusterID) error
}

type NodeSender interface {
	SendNodeMessageToNode(ctx context.Context, destinationAddress IEEEAddress, message ApplicationMessage) error
}

type EventReceiver interface {
	ReadEvent(ctx context.Context) (interface{}, error)
}

type Provider interface {
	NetworkJoining
	AdapterInfo
	NodeQuerier
	NodeBinder
	NodeSender
	EventReceiver
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
	SourceIEEEAddress    IEEEAddress
	SourceNetworkAddress NetworkAddress
	Broadcast            bool
	Secure               bool
	LinkQuality          uint8
	Sequence             uint8
	ApplicationMessage   ApplicationMessage
}

type ApplicationMessage struct {
	ClusterID           ClusterID
	SourceEndpoint      Endpoint
	DestinationEndpoint Endpoint
	Data                []byte
}

type NodeIncomingMessageEvent struct {
	Node
	IncomingMessage
}
