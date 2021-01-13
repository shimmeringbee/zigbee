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

type NodeRemover interface {
	RemoveNode(ctx context.Context, networkAddress NetworkAddress) error
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
	SendApplicationMessageToNode(ctx context.Context, destinationAddress IEEEAddress, message ApplicationMessage, requireAck bool) error
}

type EventReceiver interface {
	ReadEvent(ctx context.Context) (interface{}, error)
}

type EndpointRegistration interface {
	RegisterAdapterEndpoint(ctx context.Context, endpoint Endpoint, appProfileId ProfileID, appDeviceId uint16, appDeviceVersion uint8, inClusters []ClusterID, outClusters []ClusterID) error
}

type Provider interface {
	NetworkJoining
	AdapterInfo
	NodeQuerier
	NodeBinder
	NodeSender
	NodeRemover
	EventReceiver
	EndpointRegistration
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

type SourceAddress struct {
	IEEEAddress    IEEEAddress
	NetworkAddress NetworkAddress
}

type IncomingMessage struct {
	GroupID            GroupID
	SourceAddress      SourceAddress
	Broadcast          bool
	Secure             bool
	LinkQuality        uint8
	Sequence           uint8
	ApplicationMessage ApplicationMessage
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
