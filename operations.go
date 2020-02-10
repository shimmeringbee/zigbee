package zigbee

import (
	"context"
)

type NodeQueryer interface {
	QueryNodeDescription(ctx context.Context, networkAddress NetworkAddress) (NodeDescription, error)
	QueryNodeEndpoints(ctx context.Context, networkAddress NetworkAddress) ([]byte, error)
	QueryNodeEndpointDescription(ctx context.Context, networkAddress NetworkAddress, endpoint byte) (EndpointDescription, error)
}

type NodeBinder interface {
	BindToNode(ctx context.Context, networkAddress NetworkAddress, sourceAddress IEEEAddress, sourceEndpoint byte, destinationAddress IEEEAddress, destinationEndpoint byte, cluster ZCLClusterID) error
	UnbindFromNode(ctx context.Context, networkAddress NetworkAddress, sourceAddress IEEEAddress, sourceEndpoint byte, destinationAddress IEEEAddress, destinationEndpoint byte, cluster ZCLClusterID) error
}

type NodeSender interface {
	SendNodeMessage(ctx context.Context, destinationAddress IEEEAddress, sourceEndpoint byte, destinationEndpoint byte, cluster ZCLClusterID, data []byte) error
}

type EventReceiver interface {
	ReadEvent(ctx context.Context) (interface{}, error)
}

type BasicDeviceEvent struct {
	NetworkAddress NetworkAddress
	IEEEAddress    IEEEAddress
}

type DeviceJoinEvent BasicDeviceEvent

type DeviceAppearEvent BasicDeviceEvent

type DeviceLeaveEvent BasicDeviceEvent

type DeviceIncomingMessageEvent struct {
	GroupID              uint16
	ClusterID            ZCLClusterID
	SourceIEEEAddress    IEEEAddress
	SourceNetworkAddress NetworkAddress
	SourceEndpoint       uint8
	DestinationEndpoint  uint8
	Broadcast            bool
	Secure               bool
	LinkQuality          uint8
	Sequence             uint8
	Data                 []byte
}
