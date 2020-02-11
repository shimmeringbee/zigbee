package zigbee

import (
	"context"
	"github.com/shimmeringbee/zigbee"
	"time"
)

type NodeQueryer interface {
	QueryNodeDescription(ctx context.Context, networkAddress IEEEAddress) (NodeDescription, error)
	QueryNodeEndpoints(ctx context.Context, networkAddress IEEEAddress) ([]byte, error)
	QueryNodeEndpointDescription(ctx context.Context, networkAddress IEEEAddress, endpoint byte) (EndpointDescription, error)
}

type NodeBinder interface {
	BindNodeToController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint byte, destinationEndpoint byte, cluster ZCLClusterID) error
	UnbindNodeFromController(ctx context.Context, nodeAddress IEEEAddress, sourceEndpoint byte, destinationEndpoint byte, cluster ZCLClusterID) error
}

type NodeSender interface {
	SendNodeMessage(ctx context.Context, destinationAddress IEEEAddress, sourceEndpoint byte, destinationEndpoint byte, cluster ZCLClusterID, data []byte) error
}

type EventReceiver interface {
	ReadEvent(ctx context.Context) (interface{}, error)
}

type Device struct {
	IEEEAddress    zigbee.IEEEAddress
	NetworkAddress zigbee.NetworkAddress
	LogicalType    zigbee.LogicalType
	LQI            uint8
	Depth          uint8
	LastDiscovered time.Time
	LastReceived   time.Time
}

type DeviceJoinEvent struct {
	Device
}

type DeviceUpdateEvent struct {
	Device
}

type DeviceLeaveEvent struct {
	Device
}

type IncomingMessage struct {
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

type DeviceIncomingMessageEvent struct {
	Device
	IncomingMessage
}
