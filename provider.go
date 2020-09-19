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
	EventReceiver
	EndpointRegistration
}

type Node struct {
	IEEEAddress    IEEEAddress    `json:"ieee_address"`
	NetworkAddress NetworkAddress `json:"network_address"`
	LogicalType    LogicalType    `json:"logical_type"`
	LQI            uint8          `json:"lqi,omitempty"`
	Depth          uint8          `json:"depth,omitempty"`
	LastDiscovered time.Time      `json:"last_discovered,omitempty"`
	LastReceived   time.Time      `json:"last_received,omitempty"`
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
